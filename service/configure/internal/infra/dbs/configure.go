package dbs

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	json "github.com/json-iterator/go"

	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type msg struct {
	EnvId uint32 `json:"envId"`
	SrvId uint32 `json:"srvId"`
}

type Configure struct {
	watchQueue *redis.PubSub
}

var (
	configureIns  *Configure
	configureOnce sync.Once
)

const (
	configureMsgChannel = "configure:watcher:channel"
)

func NewConfigure() *Configure {
	configureOnce.Do(func() {
		ctx := kratosx.MustContext(context.Background())
		configureIns = &Configure{
			watchQueue: ctx.Redis().Subscribe(ctx, configureMsgChannel),
		}
	})
	return configureIns
}

// GetConfigureByEnvAndSrv 获取指定版本的模板
func (r Configure) GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*entity.Configure, error) {
	var (
		configure = entity.Configure{}
		fs        = []string{"*"}
	)

	return &configure, ctx.DB().Select(fs).Where("env_id=? and server_id=?", envId, srvId).First(&configure).Error
}

// GetConfigure 获取指定的数据
func (r Configure) GetConfigure(ctx kratosx.Context, id uint32) (*entity.Configure, error) {
	var (
		configure = entity.Configure{}
		fs        = []string{"*"}
	)

	return &configure, ctx.DB().Select(fs).First(&configure, id).Error
}

// ListConfigure 获取列表
func (r Configure) ListConfigure(ctx kratosx.Context, req *types.ListConfigureRequest) ([]*entity.Configure, uint32, error) {
	var (
		list  []*entity.Configure
		total int64
		fs    = []string{"id", "format", "version", "description", "created_at", "updated_at"}
	)

	db := ctx.DB().Model(entity.Configure{}).Select(fs).
		Where("server_id=?", req.ServerId).
		Where("env_id=?", req.EnvId)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateConfigure 创建数据
func (r Configure) CreateConfigure(ctx kratosx.Context, configure *entity.Configure) (uint32, error) {
	return configure.Id, ctx.DB().Create(configure).Error
}

// UpdateConfigure 更新数据
func (r Configure) UpdateConfigure(ctx kratosx.Context, configure *entity.Configure) error {
	return ctx.DB().Where("id = ?", configure.Id).Updates(configure).Error
}

// DeleteConfigure 删除数据
func (r Configure) DeleteConfigure(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id=?", id).Delete(&entity.Configure{}).Error
}

// BroadcastConfigure 广播配置变更
func (r Configure) BroadcastConfigure(ctx kratosx.Context, envId uint32, srvId uint32) error {
	data := msg{EnvId: envId, SrvId: srvId}
	strMsg, _ := json.MarshalToString(data)
	return ctx.Redis().Publish(ctx, configureMsgChannel, strMsg).Err()
}

// SubscribeConfigure 广播配置变更
func (r Configure) SubscribeConfigure(f func(ctx kratosx.Context, envId uint32, srvId uint32) error) {
	ctx := kratosx.MustContext(context.Background())
	for {
		m, err := r.watchQueue.ReceiveMessage(ctx)
		if err != nil {
			ctx.Logger().Warnw("msg", "subscribe configure error", "err", err.Error())
			continue
		}

		data := msg{}
		if err := json.Unmarshal([]byte(m.Payload), &data); err != nil {
			ctx.Logger().Warnw("msg", "unmarshal subscribe configure error", "err", err.Error())
			continue
		}

		// 回调业务
		if err := f(ctx, data.EnvId, data.SrvId); err != nil {
			ctx.Logger().Warnw("msg", "handle subscribe configure error", "err", err.Error())
		}
	}
}
