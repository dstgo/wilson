package route

import "github.com/gin-gonic/gin"

// E element of M
type E struct {
	Key string
	Val any
}

func (e *E) Bool() bool {
	return e.Val.(bool)
}

func (e *E) String() string {
	return e.Val.(string)
}

func (e *E) Int() int {
	return e.Val.(int)
}

func (e *E) Float() float64 {
	return e.Val.(float64)
}

func Metas(e ...E) Meta {
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
