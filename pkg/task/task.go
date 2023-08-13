package task

import (
	"context"
	"sync"
)

func NewTask(ctx context.Context) *Task {
	ctx, cancel := context.WithCancelCause(ctx)
	return &Task{
		Ctx:    ctx,
		cancel: cancel,
		onPanic: func(err any) {
			panic(err)
		},
		jobs:    make([]Job, 0, 10),
		wait:    sync.WaitGroup{},
		mutex:   sync.Mutex{},
		running: false,
	}
}

type Job func(ctx context.Context) error

type Task struct {
	Ctx     context.Context
	cancel  context.CancelCauseFunc
	jobs    []Job
	onPanic func(err any)
	wait    sync.WaitGroup
	mutex   sync.Mutex
	once    sync.Once
	err     error
	running bool
}

func (t *Task) setRunning(running bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.running = running
}

func (t *Task) ClearJobs() {
	t.mutex.Lock()
	t.jobs = t.jobs[:0]
	t.mutex.Unlock()
}

func (t *Task) OnPanic(onPanic func(err any)) {
	t.onPanic = onPanic
}

func (t *Task) AddJobs(j ...Job) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.running {
		t.jobs = append(t.jobs, j...)
	}
}

func (t *Task) Run() error {
	if !t.running {

		t.setRunning(true)
		t.wait.Add(len(t.jobs))

		for _, job := range t.jobs {
			go func(j Job) {
				// wait count--
				defer t.wait.Done()
				// panic recovery
				defer func() {
					if err := recover(); err != nil {
						if t.onPanic != nil {
							t.onPanic(err)
						}
					}
				}()

				select {
				case <-t.Ctx.Done():
					return
				default:
					// call func
					if err := j(t.Ctx); err != nil {
						t.once.Do(func() {
							t.err = err
							if t.cancel != nil {
								t.cancel(err)
							}
						})
					}
				}
			}(job)

		}
		t.wait.Wait()
		t.cancel(nil)

		t.setRunning(false)
	}
	return t.err
}
