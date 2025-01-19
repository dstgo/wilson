package configure

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	pb "github.com/dstgo/wilson/api/gen/configure/configure/v1"
)

type source struct {
	host    string
	server  string
	token   string
	context context.Context
	request pb.ConfigureClient
	client  pb.Configure_WatchConfigureClient
}

func New(host, name, token string) config.Source {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(host))
	if err != nil {
		panic("configure connect error:" + err.Error())
	}

	return &source{
		request: pb.NewConfigureClient(conn),
		host:    host,
		server:  name,
		token:   token,
		context: context.Background(),
	}
}

func NewFromEnv() config.Source {
	host := os.Getenv("CONF_HOST")
	token := os.Getenv("CONF_TOKEN")
	name := os.Getenv("APP_NAME")
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(host))
	if err != nil {
		panic("configure connect error: " + err.Error())
	}

	return &source{
		request: pb.NewConfigureClient(conn),
		host:    host,
		server:  name,
		token:   token,
		context: context.Background(),
	}
}

// Load return the config values
func (s *source) Load() ([]*config.KeyValue, error) {
	client, err := s.request.WatchConfigure(s.context, &pb.WatchConfigureRequest{
		Server: s.server,
		Token:  s.token,
	})
	if err != nil {
		return nil, err
	}

	cf, err := client.Recv()
	if err != nil {
		return nil, err
	}

	s.client = client

	kv := &config.KeyValue{
		Key:    s.host,
		Value:  []byte(cf.Content),
		Format: cf.Format,
	}

	return []*config.KeyValue{kv}, nil
}

// Watch return the watcher
func (s *source) Watch() (config.Watcher, error) {
	return newWatcher(s)
}
