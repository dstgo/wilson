package dbs

import (
	"fmt"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Resource struct {
}

var (
	resourceIns  *Resource
	resourceOnce sync.Once
)

func NewResource() *Resource {
	resourceOnce.Do(func() {
		resourceIns = &Resource{}
	})
	return resourceIns
}

// GetResourceByKeyword 获取指定数据
func (r Resource) GetResourceByKeyword(ctx kratosx.Context, keyword string) (*entity.Resource, error) {
	var (
		resource = entity.Resource{}
		fs       = []string{"*"}
	)
	return &resource, ctx.DB().Select(fs).Preload("Servers").First(&resource, "keyword=?", keyword).Error
}

// GetResource 获取指定的数据
func (r Resource) GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error) {
	var (
		resource = entity.Resource{}
		fs       = []string{"*"}
	)
	return &resource, ctx.DB().Select(fs).Preload("Servers").First(&resource, id).Error
}

// ListResource 获取列表
func (r Resource) ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	var (
		list  []*entity.Resource
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Resource{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Tag != nil {
		db = db.Where("tag = ?", *req.Tag)
	}
	if req.Private != nil {
		db = db.Where("private = ?", *req.Private)
	}
	if req.ServerId != nil {
		var ids []uint32
		if err := ctx.DB().Select("resource_id").Model(entity.ResourceServer{}).Where("server_id=?", *req.ServerId).Scan(&ids).Error; err != nil {
			return nil, 0, err
		}
		db = db.Where("id in ? or private=false", ids)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateResource 创建数据
func (r Resource) CreateResource(ctx kratosx.Context, resource *entity.Resource) (uint32, error) {
	return resource.Id, ctx.DB().Create(resource).Error
}

// UpdateResource 更新数据
func (r Resource) UpdateResource(ctx kratosx.Context, resource *entity.Resource) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("resource_id=?", resource.Id).Delete(entity.ResourceServer{}).Error; err != nil {
			return err
		}
		return tx.Updates(resource).Error
	})
}

// DeleteResource 删除数据
func (r Resource) DeleteResource(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Resource{}, id).Error
}

// ListResourceValue 获取列表
func (r Resource) ListResourceValue(ctx kratosx.Context, req *types.ListResourceValueRequest) ([]*entity.ResourceValue, error) {
	var (
		rvs []*entity.ResourceValue
		fs  = []string{"*"}
	)

	db := ctx.DB().Model(entity.ResourceValue{}).Select(fs).Preload("Resource")
	if req.ResourceId != nil {
		db = db.Where("resource_id=?", *req.ResourceId)
	}
	if req.EnvId != nil {
		db = db.Where("env_id=?", *req.EnvId)
	}
	if req.ServerId != nil {
		var ids []uint32
		if err := ctx.DB().Select("resource_id").Model(entity.ResourceServer{}).Where("server_id=?", *req.ServerId).Scan(&ids).Error; err != nil {
			return nil, err
		}
		var ids1 []uint32
		if err := ctx.DB().Select("id").Model(entity.Resource{}).Where("private=false").Scan(&ids1).Error; err != nil {
			return nil, err
		}
		ids = append(ids, ids1...)
		db = db.Where("resource_id in ?", ids)
	}

	return rvs, db.Find(&rvs).Error
}

// UpdateResourceValues 更新数据
func (r Resource) UpdateResourceValues(ctx kratosx.Context, rvs []*entity.ResourceValue) error {
	return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&rvs).Error
}

// AllResourceField 获取指定服务的全部可用资源字段
func (r Resource) AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error) {
	var (
		ids  []uint32
		list []*entity.Resource
	)
	if err := ctx.DB().Select("resource_id").
		Model(entity.ResourceServer{}).
		Where("server_id=?", sid).
		Scan(&ids).Error; err != nil {
		return nil, err
	}

	if err := ctx.DB().Model(entity.Resource{}).
		Select("fields", "keyword").
		Where("id in ? or private=false", ids).
		Find(&list).Error; err != nil {
		return nil, err
	}

	var arr []string
	for _, item := range list {
		fields := strings.Split(item.Fields, ",")
		for _, val := range fields {
			arr = append(arr, item.Keyword+"."+val)
		}
	}
	return arr, nil
}
