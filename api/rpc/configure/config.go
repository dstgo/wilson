package configure

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	configurepb "github.com/dstgo/wilson/api/gen/configure/configure/v1"
)

func New(host, token, name string) config.Source {
	return &source{
		host:    host,
		server:  name,
		token:   token,
		context: context.Background(),
	}
}

type source struct {
	host        string
	server      string
	token       string
	context     context.Context
	client      configurepb.ConfigureClient
	watchClient configurepb.Configure_WatchConfigureClient

	once sync.Once
}

func (s *source) dial() {
	s.once.Do(func() {
		conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(s.host))
		if err != nil {
			panic("configure connect error: " + err.Error())
		}
		s.client = configurepb.NewConfigureClient(conn)
	})
}

// Load return the config values
func (s *source) Load() ([]*config.KeyValue, error) {
	s.dial()

	client, err := s.client.WatchConfigure(s.context, &configurepb.WatchConfigureRequest{
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

	s.watchClient = client

	kv := &config.KeyValue{
		Key:    s.host,
		Value:  []byte(cf.Content),
		Format: cf.Format,
	}

	return []*config.KeyValue{kv}, nil
}

// Watch return the watcher
func (s *source) Watch() (config.Watcher, error) {
	s.dial()

	return newWatcher(s)
}
