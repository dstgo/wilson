package configure

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/config"
)

type watcher struct {
	source *source
	ch     chan []*config.KeyValue
	ctx    context.Context
	cancel context.CancelFunc
}

func newWatcher(s *source) (*watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{
		source: s,
		ch:     make(chan []*config.KeyValue),
		ctx:    ctx,
		cancel: cancel,
	}

	go func() {
		for {
			cf, err := s.client.Recv()
			if err != nil {
				// 接收失败，则进行主动重连
				for {
					kvs, er := s.Load()
					if er != nil {
						time.Sleep(5 * time.Second)
						continue
					}
					w.ch <- kvs
					break
				}
			} else {
				kv := &config.KeyValue{
					Key:    s.host,
					Value:  []byte(cf.Content),
					Format: cf.Format,
				}
				w.ch <- []*config.KeyValue{kv}
			}
		}
	}()

	return w, nil
}

func (w *watcher) Next() ([]*config.KeyValue, error) {
	select {
	case res := <-w.ch:
		return res, nil
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	}
}

func (w *watcher) Stop() error {
	w.cancel()
	return nil
}
