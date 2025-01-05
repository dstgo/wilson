package service

import (
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/encoding"

	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/pkg/crypto"

	"github.com/dstgo/wilson/api/gen/errors"

	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/repository"
	"github.com/dstgo/wilson/service/configure/internal/types"
	"github.com/dstgo/wilson/service/configure/internal/utils"
)

type watcher struct {
	reply types.WatcherConfigReplyFunc
	close chan string
}

type Configure struct {
	conf       *conf.Config
	repo       repository.Configure
	server     repository.Server
	env        repository.Env
	business   repository.Business
	resource   repository.Resource
	template   repository.Template
	permission repository.Permission
	mutex      sync.Mutex
	rs         map[string][]watcher
}

func NewConfigure(
	conf *conf.Config,
	repo repository.Configure,
	server repository.Server,
	env repository.Env,
	business repository.Business,
	resource repository.Resource,
	template repository.Template,
	permission repository.Permission,
) *Configure {
	srv := &Configure{
		conf:       conf,
		repo:       repo,
		server:     server,
		env:        env,
		permission: permission,
		business:   business,
		resource:   resource,
		template:   template,
		rs:         map[string][]watcher{},
	}

	// 监听配置变更广播通知
	go repo.SubscribeConfigure(func(ctx kratosx.Context, envId uint32, srvId uint32) error {
		return srv.SendWatcher(ctx, envId, srvId)
	})

	return srv
}

// GetConfigureByEnvAndSrv 获取指定标识的配置信息
func (u *Configure) GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*entity.Configure, error) {
	if !u.permission.HasServer(ctx, envId) {
		return nil, errors.NotPermissionError()
	}
	if !u.permission.HasEnv(ctx, srvId) {
		return nil, errors.NotPermissionError()
	}

	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, envId, srvId)
	if err != nil {
		return nil, errors.GetError()
	}
	return configure, nil
}

// ListConfigure 获取分页配置信息
func (u *Configure) ListConfigure(ctx kratosx.Context, req *types.ListConfigureRequest) ([]*entity.Configure, uint32, error) {
	list, total, err := u.repo.ListConfigure(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// UpdateConfigure 更新模配置
func (u *Configure) UpdateConfigure(ctx kratosx.Context, req *entity.Configure) error {
	if !u.permission.HasServer(ctx, req.EnvId) {
		return errors.NotPermissionError()
	}
	if !u.permission.HasEnv(ctx, req.ServerId) {
		return errors.NotPermissionError()
	}

	format, content, err := u.RenderCurrentTemplate(ctx, req.ServerId, req.EnvId)
	if err != nil {
		return errors.RenderTemplateError(err.Error())
	}

	// 生成数据
	req.Version = crypto.MD5([]byte(content))
	req.Format = format
	req.Content = content

	// 查询往期配置
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, req.EnvId, req.ServerId)
	if err == nil {
		if req.Version == configure.Version {
			return errors.ConfigureVersionExistError()
		}
		req.Id = configure.Id
		if err := u.repo.UpdateConfigure(ctx, req); err != nil {
			return errors.DatabaseError(err.Error())
		}
	} else {
		if _, err := u.repo.CreateConfigure(ctx, req); err != nil {
			return errors.DatabaseError(err.Error())
		}
	}

	// 广播配置变更
	if err := u.repo.BroadcastConfigure(ctx, req.EnvId, req.ServerId); err != nil {
		return errors.BroadcastConfigureError()
	}

	return nil
}

func (u *Configure) RenderCurrentTemplate(ctx kratosx.Context, srvId, envId uint32) (string, string, error) {
	if !u.permission.HasServer(ctx, srvId) {
		return "", "", errors.NotPermissionError()
	}
	if !u.permission.HasEnv(ctx, envId) {
		return "", "", errors.NotPermissionError()
	}

	template, err := u.template.CurrentTemplate(ctx, srvId)
	if err != nil {
		return "", "", err
	}

	bvs, err := u.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		ServerId: &srvId,
		EnvId:    &envId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get business value error", "err", err.Error())
		return "", "", errors.DatabaseError()
	}
	rvs, err := u.resource.ListResourceValue(ctx, &types.ListResourceValueRequest{
		ServerId: &srvId,
		EnvId:    &envId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get resource value error", "err", err.Error())
		return "", "", errors.DatabaseError()
	}
	tpr := utils.NewTemplateParser()
	content, err := tpr.RenderTemplate(bvs, rvs, template.Format, template.Content)
	if err != nil {
		return "", "", errors.RenderTemplateError(err.Error())
	}

	return template.Format, content, nil
}

// CompareConfigure 对比配置
func (u *Configure) CompareConfigure(ctx kratosx.Context, req *types.CompareConfigureRequest) ([]*types.CompareConfigureReply, error) {
	if !u.permission.HasServer(ctx, req.ServerId) {
		return nil, errors.NotPermissionError()
	}
	if !u.permission.HasEnv(ctx, req.EnvId) {
		return nil, errors.NotPermissionError()
	}

	// 获取当前配置
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, req.EnvId, req.ServerId)

	// 解析生产中的配置
	var oldKeys []string
	otc := map[string]any{}
	if err == nil {
		oe := encoding.GetCodec(configure.Format)
		if err := oe.Unmarshal([]byte(configure.Content), &otc); err != nil {
			return nil, errors.RenderTemplateError()
		}
		for key := range otc {
			oldKeys = append(oldKeys, key)
		}
	}

	// 解析上传的模板
	var curKeys []string
	format, content, err := u.RenderCurrentTemplate(ctx, req.ServerId, req.EnvId)
	if err != nil {
		return nil, errors.RenderTemplateError(err.Error())
	}
	ctc := map[string]any{}
	ce := encoding.GetCodec(format)
	if err := ce.Unmarshal([]byte(content), &ctc); err != nil {
		return nil, errors.RenderTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*types.CompareConfigureReply

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
		reply = append(reply, &types.CompareConfigureReply{Type: "add", Key: key, Cur: val})
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
		reply = append(reply, &types.CompareConfigureReply{Type: "del", Key: key, Old: val})
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

		reply = append(reply, &types.CompareConfigureReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}

func (u *Configure) channelKey(envId, srvId uint32) string {
	return fmt.Sprintf("%v:%v", envId, srvId)
}

func (u *Configure) Watch(ctx kratosx.Context, in *types.WatcherConfigRequest, reply types.WatcherConfigReplyFunc) error {
	server, err := u.server.GetServerByKeyword(ctx, in.Server)
	if err != nil {
		return errors.ServerNotFound()
	}

	env, err := u.env.GetEnvByToken(ctx, in.Token)
	if err != nil {
		return errors.TokenAuthError()
	}

	// 获取配置
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, env.Id, server.Id)
	if err != nil {
		return errors.DatabaseError()
	}

	// 第一次链接先返回结果
	if err := reply(&types.WatcherConfigureReply{
		Content: configure.Content,
		Format:  configure.Format,
	}); err != nil {
		return errors.WatchConfigureError(err.Error())
	}

	// 注册回调监听
	closer := <-u.registerWatch(env.Id, server.Id, reply)
	return errors.WatchConfigureError(closer)
}

func (u *Configure) registerWatch(envId, srvId uint32, reply types.WatcherConfigReplyFunc) <-chan string {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(envId, srvId)
	closer := make(chan string, 1)
	u.rs[key] = append(u.rs[key], watcher{
		reply: reply,
		close: closer,
	})
	return closer
}

func (u *Configure) SendWatcher(ctx kratosx.Context, envId uint32, srvId uint32) error {
	// 获取当前配置
	cfg, err := u.repo.GetConfigureByEnvAndSrv(ctx, envId, srvId)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}

	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(envId, srvId)
	for ind := 0; ind < len(u.rs[key]); ind++ {
		item := u.rs[key][ind]
		// 发送新的配置模板
		if err := item.reply(&types.WatcherConfigureReply{
			Content: cfg.Content,
			Format:  cfg.Format,
		}); err != nil {
			// 发送失败，则关闭链接通道
			item.close <- err.Error()
			u.rs[key] = append(u.rs[key][:ind], u.rs[key][ind+1:]...)
			ind--
		}
	}
	return nil
}

// func (u *Configure) SendWatcher(in *entity.Configure) {
//	u.mutex.Lock()
//	defer u.mutex.Unlock()
//
//	key := u.channelKey(in.EnvId, in.ServerId)
//	for ind := 0; ind < len(u.rs[key]); ind++ {
//		item := u.rs[key][ind]
//		// 发送新的配置模板
//		if err := item.reply(&types.WatcherConfigureReply{
//			Content: in.Content,
//			Format:  in.Format,
//		}); err != nil {
//			// 发送失败，则关闭链接通道
//			item.close <- err.Error()
//			u.rs[key] = append(u.rs[key][:ind], u.rs[key][ind+1:]...)
//			ind--
//		}
//	}
// }
