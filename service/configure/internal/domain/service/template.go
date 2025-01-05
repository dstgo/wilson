package service

import (
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/encoding"
	"gorm.io/gorm"

	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/library/db/gormtranserror"
	"github.com/dstgo/wilson/framework/pkg/crypto"

	"github.com/dstgo/wilson/api/gen/errors"

	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/repository"
	"github.com/dstgo/wilson/service/configure/internal/types"
	"github.com/dstgo/wilson/service/configure/internal/utils"
)

type Template struct {
	conf       *conf.Config
	repo       repository.Template
	permission repository.Permission
	business   repository.Business
	resource   repository.Resource
}

func NewTemplate(
	conf *conf.Config,
	repo repository.Template,
	business repository.Business,
	resource repository.Resource,
	permission repository.Permission,
) *Template {
	return &Template{
		conf:       conf,
		repo:       repo,
		business:   business,
		resource:   resource,
		permission: permission,
	}
}

// CurrentTemplate 获取当前版本的配置信息
func (t *Template) CurrentTemplate(ctx kratosx.Context, srvId uint32) (*entity.Template, error) {
	if !t.permission.HasServer(ctx, srvId) {
		return nil, errors.NotPermissionError()
	}
	template, err := t.repo.CurrentTemplate(ctx, srvId)
	if err != nil {
		if gormtranserror.Is(gorm.ErrRecordNotFound, err) {
			return nil, errors.ServerNotExistTemplateError()
		}
		return nil, errors.GetError(err.Error())
	}
	return template, nil
}

// GetTemplate 获取指定模板信息
func (t *Template) GetTemplate(ctx kratosx.Context, id uint32) (*entity.Template, error) {
	template, err := t.repo.GetTemplate(ctx, id)
	if !t.permission.HasServer(ctx, template.ServerId) {
		return nil, errors.NotPermissionError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return template, nil
}

// ListTemplate 获取分页模板信息
func (t *Template) ListTemplate(ctx kratosx.Context, req *types.ListTemplateRequest) ([]*entity.Template, uint32, error) {
	if !t.permission.HasServer(ctx, req.ServerId) {
		return nil, 0, errors.NotPermissionError()
	}
	list, total, err := t.repo.ListTemplate(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTemplate 添加模板信息
func (t *Template) CreateTemplate(ctx kratosx.Context, template *entity.Template) (uint32, error) {
	if !t.permission.HasServer(ctx, template.ServerId) {
		return 0, errors.NotPermissionError()
	}

	// 序列化数据，判断格式是否正确
	otc := map[string]any{}
	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return 0, errors.RenderTemplateError(err.Error())
	}

	// 当前的版本
	template.Version = crypto.MD5ToUpper([]byte(template.Content))

	// 查找是否存在相同的版本
	if _, err := t.repo.GetTemplateByVersion(ctx, template.Version); err == nil {
		return 0, errors.TemplateVersionExistError()
	}

	// 进行模板检查是否正确
	bsKeys, _ := t.business.AllBusinessField(ctx, template.ServerId)
	rsKeys, _ := t.resource.AllResourceField(ctx, template.ServerId)
	keys := append(bsKeys, rsKeys...)
	tp := utils.NewTemplateParser()
	if err := tp.CheckTemplate(keys, template.Content); err != nil {
		return 0, errors.RenderTemplateError(err.Error())
	}

	// 获取当前的模板
	current, err := t.repo.CurrentTemplate(ctx, template.ServerId)
	// 获取变更对比信息
	if err == nil {
		compare, err := t.CompareTemplate(ctx, &types.CompareTemplateRequest{
			Id:      current.Id,
			Format:  template.Format,
			Content: template.Content,
		})
		if err != nil {
			return 0, errors.RenderTemplateError(err.Error())
		}
		compareByte, _ := json.Marshal(compare)
		template.Compare = string(compareByte)
	}

	// 数据入库
	id, err := t.repo.CreateTemplate(ctx, template)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}

	// 使用当前配置
	_ = t.repo.UseTemplate(ctx, template.ServerId, id)
	return id, nil
}

// PreviewTemplate 预览配置
func (t *Template) PreviewTemplate(ctx kratosx.Context, req *types.PreviewTemplateRequest) (*types.PreviewTemplateReply, error) {
	if !t.permission.HasServer(ctx, req.ServerId) {
		return nil, errors.NotPermissionError()
	}
	if !t.permission.HasEnv(ctx, req.EnvId) {
		return nil, errors.NotPermissionError()
	}

	bvs, err := t.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		ServerId: &req.ServerId,
		EnvId:    &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get business value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	rvs, err := t.resource.ListResourceValue(ctx, &types.ListResourceValueRequest{
		ServerId: &req.ServerId,
		EnvId:    &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get resource value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	tp := utils.NewTemplateParser()
	content, err := tp.RenderTemplate(bvs, rvs, req.Format, req.Content)
	if err != nil {
		return nil, err
	}

	return &types.PreviewTemplateReply{
		Content: content,
		Format:  req.Format,
	}, err
}

// PreviewCurrentTemplate 预览当前配置
func (t *Template) PreviewCurrentTemplate(ctx kratosx.Context, req *types.PreviewCurrentTemplateRequest) (*types.PreviewTemplateReply, error) {
	if !t.permission.HasServer(ctx, req.ServerId) {
		return nil, errors.NotPermissionError()
	}
	if !t.permission.HasEnv(ctx, req.EnvId) {
		return nil, errors.NotPermissionError()
	}

	tp, err := t.repo.CurrentTemplate(ctx, req.ServerId)
	if err != nil {
		return nil, errors.DatabaseError(err.Error())
	}

	reply := &types.PreviewTemplateReply{
		Format: tp.Format,
	}

	bvs, err := t.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		ServerId: &req.ServerId,
		EnvId:    &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get business value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	rvs, err := t.resource.ListResourceValue(ctx, &types.ListResourceValueRequest{
		ServerId: &req.ServerId,
		EnvId:    &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get resource value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	tpr := utils.NewTemplateParser()
	tp.Content, err = tpr.RenderTemplate(bvs, rvs, tp.Format, tp.Content)
	if err != nil {
		return nil, err
	}

	return reply, err
}

// SwitchTemplate 切换指定版本信息
func (t *Template) SwitchTemplate(ctx kratosx.Context, srvId, tpId uint32) error {
	if !t.permission.HasServer(ctx, srvId) {
		return errors.NotPermissionError()
	}

	if err := t.repo.UseTemplate(ctx, srvId, tpId); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// UpdateTemplate 更新模板信息
func (t *Template) UpdateTemplate(ctx kratosx.Context, req *entity.Template) error {
	template, err := t.repo.GetTemplate(ctx, req.Id)
	if err != nil {
		return errors.UpdateError(err.Error())
	}
	if !t.permission.HasServer(ctx, template.ServerId) {
		return errors.NotPermissionError()
	}

	if err := t.repo.UpdateTemplate(ctx, req); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteUpdateTemplate 删除模板信息
func (t *Template) DeleteUpdateTemplate(ctx kratosx.Context, id uint32) error {
	template, err := t.repo.GetTemplate(ctx, id)
	if err != nil {
		return errors.UpdateError(err.Error())
	}
	if !t.permission.HasServer(ctx, template.ServerId) {
		return errors.NotPermissionError()
	}

	if err := t.repo.DeleteTemplate(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// CompareTemplate 对比变更细节
func (t *Template) CompareTemplate(ctx kratosx.Context, req *types.CompareTemplateRequest) ([]*types.CompareTemplateReply, error) {
	// 查询版本
	template, err := t.repo.GetTemplate(ctx, req.Id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	if !t.permission.HasServer(ctx, template.ServerId) {
		return nil, errors.NotPermissionError()
	}

	// 解析当前版本数据
	var oldKeys []string
	otc := map[string]any{}

	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return nil, errors.RenderTemplateError()
	}
	for key := range otc {
		oldKeys = append(oldKeys, key)
	}

	// 解析上传的模板
	var curKeys []string
	ctc := map[string]any{}

	ce := encoding.GetCodec(req.Format)
	if err := ce.Unmarshal([]byte(req.Content), &ctc); err != nil {
		return nil, errors.RenderTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*types.CompareTemplateReply

	// 新增字段
	addKeys := utils.Diff(curKeys, oldKeys)
	for _, key := range addKeys {
		var val string
		switch ctc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(ctc[key])
			val = string(b)
		default:
			val = fmt.Sprint(ctc[key])
		}
		reply = append(reply, &types.CompareTemplateReply{Type: "add", Key: key, Cur: val})
	}

	// 删除的字段
	delKeys := utils.Diff(oldKeys, curKeys)
	for _, key := range delKeys {
		var val string
		switch otc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(otc[key])
			val = string(b)
		default:
			val = fmt.Sprint(otc[key])
		}
		reply = append(reply, &types.CompareTemplateReply{Type: "del", Key: key, Old: val})
	}

	// 变更的字段
	for key, val := range ctc {
		if otc[key] == nil {
			continue
		}
		if fmt.Sprint(val) == fmt.Sprint(otc[key]) {
			continue
		}

		var oldVal string
		switch otc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(otc[key])
			oldVal = string(b)
		default:
			oldVal = fmt.Sprint(otc[key])
		}

		var curVal string
		switch ctc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(ctc[key])
			curVal = string(b)
		default:
			curVal = fmt.Sprint(ctc[key])
		}

		reply = append(reply, &types.CompareTemplateReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}
