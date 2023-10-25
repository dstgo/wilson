package task

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	ctx := context.Background()

	task, cancel := New(ctx)
	defer cancel(nil)

	w1 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w1", 2*i+1)
		}
		panic("unexpected panic")
		return nil
	}

	task.Add(w1, w1, w1, w1, w1)
	actual := task.Run()

	assert.Equal(t, nil, actual)
}

func TestCancel(t *testing.T) {
	ctx := context.Background()

	task, cancel := New(ctx)
	cancel(nil)

	w1 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w1", 2*i+1)
		}
		panic("unexpected panic")
		return nil
	}

	task.Add(w1, w1, w1, w1, w1)
	actual := task.Run()

	assert.Equal(t, context.Canceled, actual)
}

func TestPanic(t *testing.T) {
	ctx := context.Background()

	task, cancel := New(ctx)
	defer cancel(nil)

	w1 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w1", 2*i+1)
		}
		panic("unexpected panic")
		return nil
	}

	task.Add(w1, w1, w1, w1, w1)
	actual := task.Run()

	assert.NotEqual(t, nil, actual)
}

func TestTimeout(t *testing.T) {
	ctx := context.Background()

	task, cancel := WithTimeout(ctx, time.Second)
	defer cancel(nil)

	w1 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w1", 2*i+1)
		}
		return nil
	}

	task.Add(w1, w1, w1, w1, w1)
	actual := task.Run()

	assert.Equal(t, context.DeadlineExceeded, actual)
}

func TestWithDeadLine(t *testing.T) {
	ctx := context.Background()

	task, cancel := WithDeadLine(ctx, time.Now().Add(time.Second))
	defer cancel(nil)

	w1 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w1", 2*i+1)
		}
		return nil
	}

	task.Add(w1, w1, w1, w1, w1)
	actual := task.Run()

	assert.Equal(t, context.DeadlineExceeded, actual)
}

func TestClear(t *testing.T) {
	ctx := context.Background()

	task, cancel := New(ctx)
	defer cancel(nil)

	w1 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w1", 2*i+1)
		}
		return nil
	}

	w2 := func(ctx context.Context) error {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("w2", 3*i+1)
		}
		return nil
	}

	task.Add(w1, w1, w1, func(ctx context.Context) error {
		task.Add(w2, w2, w2, w2)
		return nil
	}, w1, w1)

	actual := task.Run()

	assert.Equal(t, nil, actual)

	actual = task.Run()

	assert.Equal(t, nil, actual)
}
