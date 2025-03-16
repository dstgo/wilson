package pool

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants/v2"

	"github.com/dstgo/wilson/framework/kratosx/config"
)

type Runner interface {
	Run() error
	Ctx() context.Context
}

type Pool interface {
	Go(runner Runner) error
}

type pool struct {
	pf *ants.PoolWithFunc
}

var ins *pool

func Instance() Pool {
	return ins
}

func Init(conf *config.Pool, watcher config.Watcher) {
	if conf == nil {
		return
	}

	p, err := ants.NewPoolWithFunc(conf.Size, func(i any) {
		if run, ok := i.(Runner); ok {
			if err := run.Run(); err != nil {
				log.Context(run.Ctx()).Errorf("协程任务执行失败：%s", err.Error())
			}
		}
	},
		ants.WithExpiryDuration(conf.ExpiryDuration),
		ants.WithMaxBlockingTasks(conf.MaxBlockingTasks),
		ants.WithNonblocking(conf.Nonblocking),
		ants.WithPreAlloc(conf.PreAlloc),
	)
	if err != nil {
		panic("协程池初始化失败：" + err.Error())
	}

	ins = &pool{pf: p}

	watcher("pool.size", func(value config.Value) {
		size, err := value.Int()
		if err != nil {
			log.Errorf("watch pool.size config failed: %s", err.Error())
			return
		}
		log.Infof("watch pool.size config successfully")
		if size != 0 {
			ins.pf.Tune(int(size))
		}
	})
}

func (c *pool) Go(runner Runner) error {
	return c.pf.Invoke(runner)
}

type runner struct {
	ctx context.Context
	fn  func() error
}

func (r runner) Run() error {
	select {
	case <-r.ctx.Done():
		return r.ctx.Err()
	default:
		return r.fn()
	}
}

func (r runner) Ctx() context.Context {
	return r.ctx
}

func BgRunner(fn func() error) Runner {
	return NewRunner(context.Background(), fn)
}

func NewRunner(ctx context.Context, fn func() error) Runner {
	return &runner{fn: fn, ctx: ctx}
}
