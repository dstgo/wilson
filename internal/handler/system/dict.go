package system

import (
	"context"
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/dict"
	"github.com/dstgo/wilson/internal/types/system"
	"gorm.io/gorm"
)

func NewDictResolver(ds *data.DataSource) DictResolver {
	return DictResolver{ds: ds}
}

type DictResolver struct {
	ds *data.DataSource
}

func (d DictResolver) GetDictInfo(ctx context.Context, code string) ([]dict.DictDataInfo, error) {
	db := d.ds.ORM()
	list := []dict.DictDataInfo{}

	dt, b, err := GetDictByCode(ctx, db, code, true)
	if err != nil {
		return list, err
	} else if !b {
		return list, dict.ErrDictNotFound
	}

	for _, dd := range dt.Data {

		// if is disabled, skip
		if !dd.DeletedAt.Time.IsZero() {
			continue
		}

		value, err := dd.DataValue()
		if err != nil {
			return list, dict.ErrInvalidDictDatType.Wrap(err)
		}

		list = append(list, dict.DictDataInfo{
			Label: dd.Name,
			Key:   dd.Key,
			Value: value,
			Type:  dd.Type,
		})
	}
	return list, nil
}

func (d DictResolver) ListPageDict(ctx context.Context, option dict.DictPageOption) ([]dict.DictDetail, error) {
	list := []dict.DictDetail{}
	dictType, err := ListPageDictType(ctx, d.ds.ORM(), option)
	if err != nil {
		return list, err
	}
	for _, dt := range dictType {
		list = append(list, dict.DictDetail{
			Id:        dt.Id,
			Name:      dt.Name,
			Code:      dt.Code,
			CreatedAt: dt.CreatedAt,
			UpdatedAt: dt.UpdatedAt,
		})
	}
	return list, nil
}

func (d DictResolver) CreateDict(ctx context.Context, createOpt dict.DictSaveOption) error {
	db := d.ds.ORM()
	_, b, err := GetDictByCode(ctx, db, createOpt.Code, false)
	if err != nil {
		return err
	} else if b {
		return dict.ErrDictCodeConflict
	}
	return CreateDictType(ctx, d.ds.ORM(), createOpt.Name, createOpt.Code)
}

func (d DictResolver) UpdateDict(ctx context.Context, updateOpt dict.DictUpdateOption) error {
	db := d.ds.ORM()
	dt, b, err := GetDict(ctx, db, updateOpt.Id, false)
	if err != nil {
		return err
	} else if !b {
		return dict.ErrDictNotFound
	}
	return UpdateDictType(ctx, db, dt.Id, updateOpt.Name, updateOpt.Code)
}

func (d DictResolver) RemoveDict(ctx context.Context, id uint) error {
	db := d.ds.ORM()
	dt, b, err := GetDict(ctx, db, id, false)
	if err != nil {
		return err
	} else if !b {
		return dict.ErrDictNotFound
	}
	return RemoveDictType(ctx, db, dt.Id)
}

func (d DictResolver) ListPageDictData(ctx context.Context, option dict.DictDataPageOption) ([]dict.DictDataDetail, error) {
	db := d.ds.ORM()
	list := []dict.DictDataDetail{}
	dictData, err := ListPageDictData(ctx, db, option)
	if err != nil {
		return list, err
	}
	for _, dd := range dictData {
		list = append(list, dict.DictDataDetail{
			Id:        dd.Id,
			Label:     dd.Name,
			Key:       dd.Key,
			Value:     dd.Value,
			Type:      dd.Type,
			Order:     dd.Order,
			Enable:    !dd.DeletedAt.Time.IsZero(),
			CreatedAt: dd.CreatedAt,
			UpdatedAt: dd.UpdatedAt,
		})
	}
	return list, nil
}

func (d DictResolver) CreateDictData(ctx context.Context, option dict.DictDataSaveOption) error {
	db := d.ds.ORM()
	dt, b, err := GetDict(ctx, db, option.DictId, false)
	if err != nil {
		return err
	} else if !b {
		return dict.ErrDictNotFound
	}

	_, b, err = GetDictDataByKey(ctx, db, dt.Code, option.Key)
	if err != nil {
		return err
	} else if b {
		return dict.ErrDictDataKeyConflict
	}

	return CreateDictData(ctx, db, entity.DictData{
		DictId: dt.Id,
		Name:   option.Name,
		Key:    option.Key,
		Value:  option.Value,
		Type:   option.Type,
		Order:  option.Order,
	})
}

func (d DictResolver) UpdateDictData(ctx context.Context, option dict.DictDataUpdateOption) error {
	db := d.ds.ORM()
	dd, b, err := GetDictData(ctx, db, option.Id)
	if err != nil {
		return err
	} else if !b {
		return dict.ErrDictDataNotFound
	}

	err = UpdateDictData(ctx, db, entity.DictData{
		Id:    dd.Id,
		Name:  option.Name,
		Key:   option.Key,
		Value: option.Value,
		Type:  option.Type,
		Order: option.Order,
	})

	if err != nil {
		return err
	}

	// if disabled
	if !option.Enable {
		err := RemoveDictData(ctx, db, dd.Id, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d DictResolver) RemoveDictData(ctx context.Context, id uint) error {
	db := d.ds.ORM()
	dd, b, err := GetDictData(ctx, db, id)
	if err != nil {
		return err
	} else if !b {
		return dict.ErrDictDataNotFound
	}
	return RemoveDictData(ctx, db, dd.Id, true)
}

func GetDict(ctx context.Context, db *gorm.DB, id uint, preload bool) (entity.Dict, bool, error) {
	var dictType entity.Dict
	tx := db.WithContext(ctx).Unscoped()
	if preload {
		tx = tx.Preload("Data")
	}
	result := tx.Find(&dictType, "id = ?", id)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return dictType, false, system.ErrDatabase.Wrap(err)
	}
	return dictType, found, nil
}

func GetDictByCode(ctx context.Context, db *gorm.DB, code string, preload bool) (entity.Dict, bool, error) {
	var dictType entity.Dict
	tx := db.WithContext(ctx).Unscoped()
	if preload {
		tx = tx.Preload("Data")
	}
	result := tx.Find(&dictType, "code = ?", code)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return dictType, false, system.ErrDatabase.Wrap(err)
	}
	return dictType, found, nil
}

func ListPageDictType(ctx context.Context, db *gorm.DB, option dict.DictPageOption) ([]entity.Dict, error) {
	var types []entity.Dict

	db = db.WithContext(ctx).Unscoped()
	if len(option.Search) > 0 {
		db = db.Where("name LIKE ? OR code LIKE ?", data.Like(option.Search), data.Like(option.Search))
	}
	err := db.Scopes(data.Pages(option.Page, option.Size)).Find(&types).Error
	return types, err
}

func CreateDictType(ctx context.Context, db *gorm.DB, name string, code string) error {
	err := db.WithContext(ctx).Create(&entity.Dict{
		Name: name,
		Code: code,
	}).Error

	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func UpdateDictType(ctx context.Context, db *gorm.DB, id uint, name string, code string) error {
	err := db.WithContext(ctx).Updates(&entity.Dict{
		Id:   id,
		Name: name,
		Code: code,
	}).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func RemoveDictType(ctx context.Context, db *gorm.DB, id uint) error {
	err := db.WithContext(ctx).Delete(&entity.Dict{}, "id = ?", id).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return err
}

func GetDictDataByKey(ctx context.Context, db *gorm.DB, code string, key string) (entity.DictData, bool, error) {
	dt, b, err := GetDictByCode(ctx, db, code, false)
	if err != nil {
		return entity.DictData{}, false, err
	} else if !b {
		return entity.DictData{}, false, nil
	}

	tx := db.WithContext(ctx)

	var dd entity.DictData

	err = tx.Unscoped().Model(&dt).Association("Data").Find(&dd, "key = ?", key)
	if errors.Is(err, gorm.ErrRecordNotFound) || dd.Id == 0 {
		return entity.DictData{}, false, nil
	} else if err != nil {
		return entity.DictData{}, false, err
	}

	return dd, true, nil
}

func GetDictData(ctx context.Context, db *gorm.DB, dId uint) (entity.DictData, bool, error) {

	tx := db.WithContext(ctx)

	var dd entity.DictData

	result := tx.Unscoped().Find(&dd, "id = ?", dId)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return dd, false, err
	} else if !found {
		return dd, false, nil
	}

	return dd, true, nil
}

func CreateDictData(ctx context.Context, db *gorm.DB, dictData entity.DictData) error {
	err := db.WithContext(ctx).Create(&dictData).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func UpdateDictData(ctx context.Context, db *gorm.DB, dictData entity.DictData) error {
	err := db.WithContext(ctx).Updates(&dictData).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func ListPageDictData(ctx context.Context, db *gorm.DB, option dict.DictDataPageOption) ([]entity.DictData, error) {
	d, b, err := GetDictByCode(ctx, db, option.Code, false)
	if err != nil {
		return nil, err
	} else if !b {
		return nil, dict.ErrDictNotFound
	}

	var dds []entity.DictData

	tx := db.WithContext(ctx)
	tx = tx.Unscoped().Model(&d).Scopes(data.Pages(option.Page, option.Size))
	if option.Search != "" {
		like := data.Like(option.Search)
		tx = tx.Where("name = ? OR key LIKE ? OR value LIKE ?", like, like, like)
	}
	err = tx.Association("Data").Find(&dds)
	if err != nil {
		return nil, system.ErrDatabase.Wrap(err)
	}
	return dds, nil
}

func RemoveDictData(ctx context.Context, db *gorm.DB, ddId uint, force bool) error {
	db = db.WithContext(ctx)
	if force {
		db = db.Unscoped()
	}
	err := db.Delete(entity.DictData{}, "id = ?", ddId).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}
