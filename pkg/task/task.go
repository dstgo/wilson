package task

import (
	"context"
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
	"sync"
	"time"
)

type Task interface {
	Add(j ...Worker)
	Run() error
	OnPanic(onPanic func(err any))
	Ctx() context.Context
	Cancel(err error)
	Clear() error
}

const DefaultJobSize = 10

var ErrorTaskStillRunning = errors.New("task still running")

var ErrJobPanic = errors.New("job panic")

var t Task = &task{}

func New(ctx context.Context) (Task, context.CancelCauseFunc) {
	ctx, cancel := context.WithCancel(ctx)
	task := newTask(ctx, cancel)
	return task, task.Cancel
}

func WithTimeout(ctx context.Context, timeout time.Duration) (Task, context.CancelCauseFunc) {
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	task := newTask(timeoutCtx, cancel)
	return task, task.Cancel
}

func WithDeadLine(ctx context.Context, deadline time.Time) (Task, context.CancelCauseFunc) {
	withDeadline, cancelFunc := context.WithDeadline(ctx, deadline)
	task := newTask(withDeadline, cancelFunc)
	return task, task.Cancel
}

func newTask(ctx context.Context, cancelFunc context.CancelFunc) *task {
	return &task{
		ctx:     ctx,
		cancel:  cancelFunc,
		jobs:    make([]Worker, 0, DefaultJobSize),
		wait:    sync.WaitGroup{},
		mutex:   sync.Mutex{},
		running: false,
	}
}

type Worker func(ctx context.Context) error

func (w Worker) work(ctx context.Context) error {
	return w(ctx)
}

type task struct {
	mutex sync.Mutex
	wait  sync.WaitGroup
	once  sync.Once

	ctx    context.Context
	cancel context.CancelFunc

	jobs    []Worker
	waitJos []Worker
	onPanic func(err any)

	err     error
	running bool
}

func (t *task) Ctx() context.Context {
	return t.ctx
}

func (t *task) Cancel(err error) {
	t.once.Do(func() {
		t.err = err
		if t.cancel != nil {
			t.cancel()
		}
	})
}

func (t *task) setRunning(running bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.running = running
}

func (t *task) Clear() error {
	if t.running {
		return ErrorTaskStillRunning
	}

	t.wait.Wait()
	t.jobs = append(t.waitJos)
	t.waitJos = t.waitJos[:0]

	return nil
}

func (t *task) OnPanic(onPanic func(err any)) {
	t.onPanic = onPanic
}

func (t *task) Add(j ...Worker) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.running {
		t.jobs = append(t.jobs, j...)
	} else {
		t.waitJos = append(t.waitJos, j...)
	}
}

// Run workers until all of them finished
// if finished all jobs successfully, it will return nil
// else if ctx was cancelled, it will return context.Canceled
// else it will return the error that worker returned
func (t *task) Run() error {
	if !t.running {

		t.setRunning(true)
		t.wait.Add(len(t.jobs))

		for _, job := range t.jobs {
			go func(w Worker) {
				// wait count--
				defer t.wait.Done()

				taskCh := make(chan struct{})
				defer close(taskCh)

				// worker goroutine
				go func() {
					// panic recovery
					defer func() {
						if err := recover(); err != nil {
							if t.onPanic != nil {
								t.onPanic(err)
							}
							t.Cancel(errors2.Wrap(ErrJobPanic, fmt.Sprintf("%+v", err)))
						}
					}()

					// call func
					if err := w.work(t.ctx); err != nil {
						t.Cancel(err)
					}

					taskCh <- struct{}{}
				}()

				select {
				case <-t.ctx.Done():
					return
				case <-taskCh:
					return
				}
			}(job)
		}

		t.wait.Wait()
		t.setRunning(false)
		t.Clear()
	}

	// if all jobs finished normally, t.err should be nil
	if t.err == nil {
		// may be ctx was canceled
		t.err = t.ctx.Err()
	}

	return t.err
}
