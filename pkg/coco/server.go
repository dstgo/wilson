package coco

import (
	"context"
	"github.com/dstgo/task"
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Core
// coco Core of http server
type Core struct {

	// logger interface
	logger *logrus.Logger

	// context
	ctx context.Context
	// gin engine
	engine *gin.Engine
	// server http.Server
	server *http.Server
	// router
	router *route.Router
	// global config variable
	config any
}

func (c *Core) L() *logrus.Logger {
	return c.logger
}

func (c *Core) Engine() *gin.Engine {
	return c.engine
}

func (c *Core) Server() *http.Server {
	return c.server
}

func (c *Core) Ctx() context.Context {
	return c.ctx
}

func (c *Core) Cfg() any {
	return c.config
}

func (c *Core) RootRouter() *route.Router {
	return c.router
}

// Coco
// coco http server
type Coco struct {
	// coco Core
	c *Core

	signals []os.Signal
	// sync control
	mu     sync.Mutex
	once   sync.Once
	listen chan struct{}

	// interrupt tasks
	interrupt []InterruptFn

	// components tasks, sync tasks run first, then async tasks
	preSync  []ComponentFn
	preAsync []ComponentFn

	postAsync []ComponentFn

	OnPanic func(err any)
	cancel  context.CancelCauseFunc
}

func (c *Coco) Core() *Core {
	return c.c
}

func (c *Coco) SetSignals(s ...os.Signal) {
	c.signals = s
}

func (c *Coco) AddPreSyncCs(cs ...ComponentFn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.preSync = append(c.preSync, cs...)
}

func (c *Coco) AddPreAsyncCs(cs ...ComponentFn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.preAsync = append(c.preAsync, cs...)
}

func (c *Coco) AddPostAsyncCs(cs ...ComponentFn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.postAsync = append(c.postAsync, cs...)
}

func (c *Coco) OnInterrupt(fn ...InterruptFn) {
	c.interrupt = append(c.interrupt, fn...)
}

// onInterrupt
// param s ...os.Signal
// run interrupt fn when received expected os signal
func (c *Coco) onInterrupt(s ...os.Signal) {
	ch := make(chan os.Signal)
	signal.Notify(ch, s...)
	sig := <-ch
	c.cancel(errors.New(sig.String()))

	for _, fn := range c.interrupt {
		fn(c.c, sig)
	}
	os.Exit(1)
}

// execPreComponents
// pre components run before server starting to listen
// usually they are some initial tasks
// async run after sync
func (c *Coco) execPreComponents() {
	// run sync components
	for _, fn := range c.preSync {
		fn(c.c)
	}

	preTask := task.NewTask(func(err any) {
		if c.OnPanic != nil {
			c.OnPanic(err)
		}
	})

	// async components
	for _, fn := range c.preAsync {
		fn := fn
		preTask.AddJobs(func() {
			fn(c.c)
		})
	}

	preTask.Run()
}

// execPostComponents
// post components run after server listening
func (c *Coco) execPostComponents(ch chan struct{}) {
	// blocked unless http server started listening
	<-ch

	postTask := task.NewTask(func(err any) {
		if c.OnPanic != nil {
			c.OnPanic(err)
		}
	})

	for _, fn := range c.postAsync {
		fn := fn
		postTask.AddJobs(func() {
			fn(c.c)
		})
	}

	postTask.Run()
}

func (c *Coco) prepare() {
	// new goroutine listening os signal
	go c.onInterrupt(c.signals...)

	c.execPreComponents()

	// async post tasks
	go c.execPostComponents(c.listen)
}
func (c *Coco) serve() {
	c.once.Do(func() {
		c.prepare()
		c.listen <- struct{}{}
	})
}

func (c *Coco) Run() error {
	c.serve()
	return c.c.server.ListenAndServe()
}

func (c *Coco) RunTLS(certFile, keyFile string) error {
	c.serve()
	return c.c.server.ListenAndServeTLS(certFile, keyFile)
}

func (c *Coco) RunListener(l net.Listener) error {
	c.serve()
	return c.c.server.Serve(l)
}

func (c *Coco) RunListenerTls(l net.Listener, certFile, keyFile string) error {
	c.serve()
	return c.c.server.ServeTLS(l, certFile, keyFile)
}

func newCore(opts ...ComponentFn) *Core {
	c := Core{}

	// apply options
	for _, opt := range opts {
		opt(&c)
	}

	if c.config == nil {
		c.config = DefaultConfig()
	}

	if c.ctx == nil {
		c.ctx = context.Background()
	}

	if c.logger == nil {
		c.logger = logrus.New()
	}

	if c.engine == nil {
		c.engine = gin.Default()
	}

	if c.server == nil {
		c.server = &http.Server{}
		c.server.Addr = ":8080"
	}

	gin.DisableBindValidation()
	// attach gin handler
	c.server.Handler = c.engine

	return &c
}

// New create new coco with options
// param opts ...ComponentFn
// return *Coco
func New(ctx context.Context, opts ...ComponentFn) *Coco {
	core := newCore(opts...)
	return NewWithCore(ctx, core)
}

// NewWithCore
// param Core *Core
// return *Coco
// with custom Core
func NewWithCore(ctx context.Context, core *Core) *Coco {
	cancelCtx, cancel := context.WithCancelCause(ctx)
	core.ctx = cancelCtx
	return &Coco{
		cancel:  cancel,
		c:       core,
		signals: []os.Signal{syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGTERM},
		listen:  make(chan struct{}),
	}
}
