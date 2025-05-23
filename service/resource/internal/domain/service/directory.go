package service

import (
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/pkg/tree"

	"github.com/dstgo/wilson/api/gen/errors"
	"github.com/dstgo/wilson/service/resource/internal/conf"
	"github.com/dstgo/wilson/service/resource/internal/domain/entity"
	"github.com/dstgo/wilson/service/resource/internal/domain/repository"
	"github.com/dstgo/wilson/service/resource/internal/types"
)

type Directory struct {
	conf *conf.Config
	repo repository.Directory
}

func NewDirectory(
	conf *conf.Config,
	repo repository.Directory,
) *Directory {
	return &Directory{
		conf: conf,
		repo: repo,
	}
}

// GetDirectory 获取指定的文件目录信息
func (u *Directory) GetDirectory(ctx kratosx.Context, id uint32) (*entity.Directory, error) {
	res, err := u.repo.GetDirectory(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get directory error", "err", err.Error())
		return nil, errors.GetErrorWrap(err)
	}
	return res, nil
}

// ListDirectory 获取文件目录信息列表树
func (u *Directory) ListDirectory(ctx kratosx.Context, req *types.ListDirectoryRequest) ([]*entity.Directory, uint32, error) {
	list, total, err := u.repo.ListDirectory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list directory error", "err", err.Error())
		return nil, 0, errors.ListErrorWrap(err)
	}
	return tree.BuildArrayTree(list), total, nil
}

// CreateDirectory 创建文件目录信息
func (u *Directory) CreateDirectory(ctx kratosx.Context, req *entity.Directory) (uint32, error) {
	id, err := u.repo.CreateDirectory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create directory error", "err", err.Error())
		return 0, errors.CreateErrorWrap(err)
	}
	return id, nil
}

// UpdateDirectory 更新文件目录信息
func (u *Directory) UpdateDirectory(ctx kratosx.Context, req *entity.Directory) error {
	if err := u.repo.UpdateDirectory(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update directory error", "err", err.Error())
		return errors.UpdateErrorWrap(err)
	}
	return nil
}

// DeleteDirectory 删除文件目录信息
func (u *Directory) DeleteDirectory(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteDirectory(ctx, ids)
	if err != nil {
		ctx.Logger().Warnw("msg", "delete directory error", "err", err.Error())
		return 0, errors.DeleteErrorWrap(err)
	}
	return total, nil
}
