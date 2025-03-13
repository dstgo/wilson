package middleware

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
)

func Recovery() middleware.Middleware {
	handler := func(ctx context.Context, req, err any) error {
		e, ok := err.(*errors.Error)
		if ok {
			return e
		}
		return errors.InternalServer("internal server error", fmt.Sprintf("%+v", err))
	}

	return recovery.Recovery(recovery.WithHandler(handler))
}
