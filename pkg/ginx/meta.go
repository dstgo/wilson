package ginx

import "github.com/gin-gonic/gin"

// E element of M
type E struct {
	Key string
	Val any
}

func (e *E) Bool() bool {
	if b, ok := e.Val.(bool); ok {
		return b
	}
	return false
}

func (e *E) String() string {
	if b, ok := e.Val.(string); ok {
		return b
	}
	return ""
}

func (e *E) Int() int {
	if b, ok := e.Val.(int); ok {
		return b
	}
	return 0
}

func (e *E) Float() float64 {
	if b, ok := e.Val.(float64); ok {
		return b
	}
	return 0
}

func M(e ...E) Meta {
	if len(e) == 0 {
		return nil
	}
	meta := make(Meta)
	for _, E := range e {
		meta[E.Key] = E.Val
	}
	return meta
}

// Meta
// route meta info
type Meta map[string]any

func (m Meta) Get(key string) (E, bool) {
	v, e := m[key]
	if !e {
		return E{}, false
	}
	return E{Key: key, Val: v}, true
}

func (m Meta) Has(key string) bool {
	_, b := m.Get(key)
	return b
}

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
