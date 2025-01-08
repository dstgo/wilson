package ip

import "context"

type CtxKey struct{}

func ClientIP(ctx context.Context) string {
	ip, _ := ctx.Value(CtxKey{}).(string)
	return ip
}
