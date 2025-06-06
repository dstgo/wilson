package app

import (
	"context"
	"io"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx"
	ktypes "github.com/dstgo/wilson/framework/kratosx/types"
	"github.com/dstgo/wilson/framework/pkg/valx"

	"github.com/dstgo/wilson/api/gen/errors"
	pb "github.com/dstgo/wilson/api/gen/resource/file/v1"
	"github.com/dstgo/wilson/service/resource/internal/conf"
	"github.com/dstgo/wilson/service/resource/internal/domain/entity"
	"github.com/dstgo/wilson/service/resource/internal/domain/service"
	"github.com/dstgo/wilson/service/resource/internal/infra/dbs"
	"github.com/dstgo/wilson/service/resource/internal/infra/store"
	"github.com/dstgo/wilson/service/resource/internal/types"
)

type File struct {
	pb.UnimplementedFileServer
	srv *service.File
}

func NewFile(conf *conf.Config) *File {
	return &File{
		srv: service.NewFile(conf, dbs.NewFile(), dbs.NewDirectory(conf), store.NewStore(conf)),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		app := NewFile(c)
		pb.RegisterFileHTTPServer(hs, app)
		pb.RegisterFileServer(gs, app)

		cr := hs.Route("/")
		cr.GET("/resource/api/v1/static/{expire}/{sign}/{src}", app.srv.SrcBlob())
		cr.POST("/resource/api/v1/upload", app.Upload())
	})
}

// GetFile 获取指定的文件信息
func (s *File) GetFile(c context.Context, req *pb.GetFileRequest) (*pb.GetFileReply, error) {
	result, err := s.srv.GetFile(kratosx.MustContext(c), &types.GetFileRequest{
		Id:  req.Id,
		Sha: req.Sha,
		Src: req.Src,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetFileReply{
		Id:          result.Id,
		DirectoryId: result.DirectoryId,
		Name:        result.Name,
		Type:        result.Type,
		Size:        result.Size,
		Sha:         result.Sha,
		Src:         result.Src,
		Url:         result.Url,
		Status:      result.Status,
		UploadId:    result.UploadId,
		ChunkCount:  result.ChunkCount,
		CreatedAt:   uint32(result.CreatedAt),
		UpdatedAt:   uint32(result.UpdatedAt),
	}, nil
}

// ListFile 获取文件信息列表
func (s *File) ListFile(c context.Context, req *pb.ListFileRequest) (*pb.ListFileReply, error) {
	list, total, err := s.srv.ListFile(kratosx.MustContext(c), &types.ListFileRequest{
		Page:        req.Page,
		PageSize:    req.PageSize,
		Order:       req.Order,
		OrderBy:     req.OrderBy,
		DirectoryId: req.DirectoryId,
		Status:      req.Status,
		Name:        req.Name,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListFileReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListFileReply_File{
			Id:          item.Id,
			DirectoryId: item.DirectoryId,
			Name:        item.Name,
			Type:        item.Type,
			Size:        item.Size,
			Sha:         item.Sha,
			Src:         item.Src,
			Url:         item.Url,
			Status:      item.Status,
			UploadId:    item.UploadId,
			ChunkCount:  item.ChunkCount,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// PrepareUploadFile 预上传文件信息
func (s *File) PrepareUploadFile(c context.Context, req *pb.PrepareUploadFileRequest) (*pb.PrepareUploadFileReply, error) {
	if req.DirectoryPath == nil && req.DirectoryId == nil {
		return nil, errors.ParamsError()
	}

	res, err := s.srv.PrepareUploadFile(kratosx.MustContext(c), &types.PrepareUploadFileRequest{
		DirectoryId:   req.DirectoryId,
		DirectoryPath: req.DirectoryPath,
		Name:          req.Name,
		Size:          req.Size,
		Sha:           req.Sha,
	})
	if err != nil {
		return nil, err
	}

	return &pb.PrepareUploadFileReply{
		Uploaded:     res.Uploaded,
		Src:          res.Src,
		ChunkSize:    res.ChunkSize,
		ChunkCount:   res.ChunkCount,
		UploadId:     res.UploadId,
		UploadChunks: res.UploadChunks,
		Sha:          res.Sha,
		Url:          res.Url,
	}, nil
}

// UploadFile 上传文件信息
func (s *File) UploadFile(c context.Context, req *pb.UploadFileRequest) (*pb.UploadFileReply, error) {
	reply, err := s.srv.UploadFile(kratosx.MustContext(c), &types.UploadFileRequest{
		UploadId: req.UploadId,
		Index:    req.Index,
		Data:     req.Data,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UploadFileReply{
		Src: reply.Src,
		Sha: reply.Sha,
		Url: reply.Url,
	}, nil
}

// UpdateFile 更新文件信息
func (s *File) UpdateFile(c context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileReply, error) {
	if err := s.srv.UpdateFile(kratosx.MustContext(c), &entity.File{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		DirectoryId: req.DirectoryId,
		Name:        req.Name,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateFileReply{}, nil
}

// DeleteFile 删除文件信息
func (s *File) DeleteFile(c context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileReply, error) {
	total, err := s.srv.DeleteFile(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFileReply{Total: total}, nil
}

func (s *File) Upload() http.HandlerFunc {
	return func(ctx http.Context) error {
		var in pb.UploadFileRequest

		in.UploadId = ctx.Request().FormValue("uploadId")
		in.Index = valx.StrToUint32(ctx.Request().FormValue("index"))
		file, _, err := ctx.Request().FormFile("data")
		if err != nil {
			return errors.UploadFileErrorWrap(err)
		}

		in.Data, err = io.ReadAll(file)
		if err != nil {
			return errors.UploadFileErrorWrap(err)
		}
		if in.UploadId == "" || in.Index == 0 || len(in.Data) == 0 {
			return errors.ParamsError()
		}

		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			upReq, ok := req.(*pb.UploadFileRequest)
			if !ok {
				return nil, errors.TransformErrorf("invalid request type, expected *pb.UploadFileRequest, got %T", req)
			}
			return s.UploadFile(ctx, upReq)
		})

		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply, ok := out.(*pb.UploadFileReply)
		if !ok {
			return errors.TransformErrorf("invalid reply type, expected *pb.UploadFileReply, got %T", out)
		}
		return ctx.Result(200, reply)
	}
}
