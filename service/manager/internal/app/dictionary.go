package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx"
	ktypes "github.com/dstgo/wilson/framework/kratosx/types"
	"github.com/dstgo/wilson/framework/pkg/valx"

	"github.com/dstgo/wilson/api/gen/errors"
	pb "github.com/dstgo/wilson/api/gen/manager/dictionary/v1"
	"github.com/dstgo/wilson/service/manager/internal/conf"
	"github.com/dstgo/wilson/service/manager/internal/domain/entity"
	"github.com/dstgo/wilson/service/manager/internal/domain/service"
	"github.com/dstgo/wilson/service/manager/internal/infra/dbs"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type Dictionary struct {
	pb.UnimplementedDictionaryServer
	srv *service.Dictionary
}

func NewDictionary(conf *conf.Config) *Dictionary {
	return &Dictionary{
		srv: service.NewDictionary(conf, dbs.NewDictionary()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewDictionary(c)
		pb.RegisterDictionaryHTTPServer(hs, srv)
		pb.RegisterDictionaryServer(gs, srv)
	})
}

// GetDictionary 获取指定的字典目录
func (s *Dictionary) GetDictionary(c context.Context, req *pb.GetDictionaryRequest) (*pb.GetDictionaryReply, error) {
	result, err := s.srv.GetDictionary(kratosx.MustContext(c), &types.GetDictionaryRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetDictionaryReply{
		Id:          result.Id,
		Keyword:     result.Keyword,
		Name:        result.Name,
		Description: result.Description,
		CreatedAt:   uint32(result.CreatedAt),
		UpdatedAt:   uint32(result.UpdatedAt),
	}, nil
}

// ListDictionary 获取字典目录列表
func (s *Dictionary) ListDictionary(c context.Context, req *pb.ListDictionaryRequest) (*pb.ListDictionaryReply, error) {
	ctx := kratosx.MustContext(c)
	list, total, err := s.srv.ListDictionary(ctx, &types.ListDictionaryRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListDictionaryReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListDictionaryReply_Dictionary{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateDictionary 创建字典目录
func (s *Dictionary) CreateDictionary(c context.Context, req *pb.CreateDictionaryRequest) (*pb.CreateDictionaryReply, error) {
	id, err := s.srv.CreateDictionary(kratosx.MustContext(c), &entity.Dictionary{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDictionaryReply{Id: id}, nil
}

// UpdateDictionary 更新字典目录
func (s *Dictionary) UpdateDictionary(c context.Context, req *pb.UpdateDictionaryRequest) (*pb.UpdateDictionaryReply, error) {
	if err := s.srv.UpdateDictionary(kratosx.MustContext(c), &entity.Dictionary{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateDictionaryReply{}, nil
}

// DeleteDictionary 删除字典目录
func (s *Dictionary) DeleteDictionary(c context.Context, req *pb.DeleteDictionaryRequest) (*pb.DeleteDictionaryReply, error) {
	return &pb.DeleteDictionaryReply{}, s.srv.DeleteDictionary(kratosx.MustContext(c), req.Id)
}

// ListDictionaryValue 获取字典值目录列表
func (s *Dictionary) ListDictionaryValue(c context.Context, req *pb.ListDictionaryValueRequest) (*pb.ListDictionaryValueReply, error) {
	ctx := kratosx.MustContext(c)
	result, total, err := s.srv.ListDictionaryValue(ctx, &types.ListDictionaryValueRequest{
		Page:         req.Page,
		PageSize:     req.PageSize,
		DictionaryId: req.DictionaryId,
		Label:        req.Label,
		Value:        req.Value,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListDictionaryValueReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateDictionaryValue 创建字典值目录
func (s *Dictionary) CreateDictionaryValue(c context.Context, req *pb.CreateDictionaryValueRequest) (*pb.CreateDictionaryValueReply, error) {
	id, err := s.srv.CreateDictionaryValue(kratosx.MustContext(c), &entity.DictionaryValue{
		DictionaryId: req.DictionaryId,
		Label:        req.Label,
		Value:        req.Value,
		Status:       req.Status,
		Weight:       req.Weight,
		Type:         req.Type,
		Extra:        req.Extra,
		Description:  req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDictionaryValueReply{Id: id}, nil
}

// UpdateDictionaryValue 更新字典值目录
func (s *Dictionary) UpdateDictionaryValue(c context.Context, req *pb.UpdateDictionaryValueRequest) (*pb.UpdateDictionaryValueReply, error) {
	if err := s.srv.UpdateDictionaryValue(kratosx.MustContext(c), &entity.DictionaryValue{
		BaseModel:    ktypes.BaseModel{Id: req.Id},
		DictionaryId: req.DictionaryId,
		Label:        req.Label,
		Value:        req.Value,
		Weight:       req.Weight,
		Type:         req.Type,
		Extra:        req.Extra,
		Description:  req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateDictionaryValueReply{}, nil
}

// UpdateDictionaryValueStatus 更新字典值目录状态
func (s *Dictionary) UpdateDictionaryValueStatus(c context.Context, req *pb.UpdateDictionaryValueStatusRequest) (*pb.UpdateDictionaryValueStatusReply, error) {
	return &pb.UpdateDictionaryValueStatusReply{}, s.srv.UpdateDictionaryValueStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// DeleteDictionaryValue 删除字典值目录
func (s *Dictionary) DeleteDictionaryValue(c context.Context, req *pb.DeleteDictionaryValueRequest) (*pb.DeleteDictionaryValueReply, error) {
	return &pb.DeleteDictionaryValueReply{}, s.srv.DeleteDictionaryValue(kratosx.MustContext(c), req.Id)
}

func (s *Dictionary) GetDictionaryValues(c context.Context, req *pb.GetDictionaryValuesRequest) (*pb.GetDictionaryValuesReply, error) {
	res, err := s.srv.GetDictionaryValues(kratosx.MustContext(c), req.Keywords)
	if err != nil {
		return nil, err
	}

	reply := pb.GetDictionaryValuesReply{Dict: make(map[string]*pb.GetDictionaryValuesReply_Value)}
	for key, values := range res {
		if reply.Dict[key] == nil {
			reply.Dict[key] = &pb.GetDictionaryValuesReply_Value{
				List: make([]*pb.GetDictionaryValuesReply_Value_Item, 0),
			}
		}
		if err := valx.Transform(values, &reply.Dict[key].List); err != nil {
			return nil, errors.TransformErrorWrap(err)
		}
	}
	return &reply, nil
}
