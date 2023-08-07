package route

import "github.com/gin-gonic/gin"

// E element of M
type E struct {
	K string
	V any
}

func (e *E) Bool() bool {
	return e.V.(bool)
}

func (e *E) String() string {
	return e.V.(string)
}

func (e *E) Int() int {
	return e.V.(int)
}

func (e *E) Float() float64 {
	return e.V.(float64)
}

func Metas(e ...E) Meta {
	if len(e) == 0 {
		return nil
	}
	meta := make(Meta)
	for _, E := range e {
		meta[E.K] = E.V
	}
	return meta
}

// Meta
// route meta info
type Meta = map[string]any

// MetaKey
// meta key, could update if you need
var MetaKey = "coco.route.meta"

func MetaHandler(meta Meta) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if meta == nil {
			meta = make(Meta)
		}
		ctx.Set(MetaKey, meta)
	}
}
func MetaFromCtx(ctx *gin.Context) Meta {
	value, e := ctx.Get(MetaKey)
	if !e {
		value = make(Meta)
	}
	return value.(Meta)
}
