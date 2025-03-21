package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx"

	pb "github.com/dstgo/wilson/api/gen/manager/system/v1"
	"github.com/dstgo/wilson/service/manager/internal/conf"
	"github.com/dstgo/wilson/service/manager/internal/domain/service"
	"github.com/dstgo/wilson/service/manager/internal/infra/dbs"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type System struct {
	pb.UnimplementedSystemServer
	srv *service.System
}

func NewSystem(conf *conf.Config) *System {
	return &System{
		srv: service.NewSystem(conf, dbs.NewDictionary()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewSystem(c)
		pb.RegisterSystemHTTPServer(hs, srv)
		pb.RegisterSystemServer(gs, srv)
	})
}

// GetSystemSetting 获取系统设置
func (s *System) GetSystemSetting(c context.Context, _ *pb.GetSystemSettingRequest) (*pb.GetSystemSettingReply, error) {
	setting := s.srv.GetSystemSetting(kratosx.MustContext(c), &types.GetSystemSettingRequest{})

	reply := pb.GetSystemSettingReply{
		Debug:              setting.Debug,
		Title:              setting.Title,
		Desc:               setting.Desc,
		Copyright:          setting.Copyright,
		Logo:               setting.Logo,
		Watermark:          setting.Watermark,
		ChangePasswordType: setting.ChangePasswordType,
	}
	if len(setting.Dictionaries) != 0 {
		dictArr := make(map[string]*pb.GetSystemSettingReply_DictionaryValueList)
		for _, item := range setting.Dictionaries {
			if dictArr[item.Keyword] == nil {
				dictArr[item.Keyword] = &pb.GetSystemSettingReply_DictionaryValueList{}
			}
			dv := &pb.DictionaryValue{
				Label: item.Label,
				Value: item.Value,
				Type:  item.Type,
				Extra: item.Extra,
			}
			dictArr[item.Keyword].List = append(dictArr[item.Keyword].List, dv)
		}
		reply.Dictionaries = dictArr
	}
	return &reply, nil
}
