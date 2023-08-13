package shutdown

import (
	"context"
	"os"
	"os/signal"
)

type CallBack func(signal os.Signal)

func ListenSignal(fn func(signal os.Signal), sigs ...os.Signal) {
	ch := make(chan os.Signal)
	signal.Notify(ch, sigs...)
	fn(<-ch)
}

func NotifyCtx(ctx context.Context, sigs ...os.Signal) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, sigs...)
}
