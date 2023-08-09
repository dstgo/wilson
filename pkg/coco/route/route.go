package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type Walker func(info RouterInfo) error

// Routes which store all registered route info
type Routes struct {
	rs []RouterInfo
}

func (r *Routes) AddRoute(info RouterInfo) {
	r.rs = append(r.rs, info)
}

func (r *Routes) Walk(w Walker) error {
	for _, info := range r.rs {
		err := w(info)
		if err != nil {
			return err
		}
	}
	return nil
}

type RouterGroup interface {
	Use(handler ...gin.HandlerFunc)
	Group(path string, meta Meta, handler ...gin.HandlerFunc) RouterGroup
	Handle(method, path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	GET(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	POST(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	DELETE(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	PUT(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	HEAD(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	OPTIONS(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	PATCH(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo
	Any(path string, meta Meta, handler ...gin.HandlerFunc) []RouterInfo
	Match(method []string, path string, meta Meta, handler ...gin.HandlerFunc) []RouterInfo
}

// RouterInfo describe a registered info
// if IsGroup is true, the root represent itself
type RouterInfo struct {
	IsGroup bool
	Group   *gin.RouterGroup

	Method   string
	FullPath string
	Meta     Meta

	Chain gin.HandlersChain
}

type Router struct {
	root       *gin.RouterGroup
	middleware []gin.HandlerFunc
	routes     *Routes
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		root:   &engine.RouterGroup,
		routes: &Routes{rs: make([]RouterInfo, 0, 10)},
	}
}

// Walk
// iterate each Group with fn
// you should call this func on root router
// param walker Walker
// return error
func (w *Router) Walk(walker Walker) error {
	return w.routes.Walk(walker)
}

func (w *Router) addRoute(isGroup bool, method, path string, meta Meta, handlers ...gin.HandlerFunc) RouterInfo {
	// join fullpath
	fullpath := joinPaths(w.root.BasePath(), path)

	fullHandlers := append(w.middleware, handlers...)
	// meta handler
	if meta != nil {
		fullHandlers = append([]gin.HandlerFunc{MetaHandler(meta)}, fullHandlers...)
	}
	handlers = fullHandlers

	info := RouterInfo{
		Method:   method,
		FullPath: fullpath,
		Meta:     meta,
		Chain:    handlers,
		IsGroup:  isGroup,
	}

	if isGroup {
		info.Group = w.root.Group(path)
		info.IsGroup = true
	} else {
		info.Group = w.root
		w.root.Handle(method, path, handlers...)
	}

	w.routes.AddRoute(info)

	return info
}

func (w *Router) Use(handler ...gin.HandlerFunc) {
	w.middleware = append(w.middleware, handler...)
}

func (w *Router) Attach(handler ...gin.HandlerFunc) {
	w.root.Use(handler...)
}

func (w *Router) Group(path string, meta Meta, handler ...gin.HandlerFunc) RouterGroup {
	info := w.addRoute(true, "", path, meta, handler...)
	return &Router{
		root:       info.Group,
		middleware: append(w.middleware, handler...),
		routes:     w.routes,
	}
}

func (w *Router) Handle(method, path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.addRoute(false, method, path, meta, handler...)
}

func (w *Router) GET(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodGet, path, meta, handler...)
}

func (w *Router) POST(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodPost, path, meta, handler...)
}

func (w *Router) DELETE(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodDelete, path, meta, handler...)
}

func (w *Router) PUT(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodPut, path, meta, handler...)
}

func (w *Router) HEAD(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodHead, path, meta, handler...)
}

func (w *Router) OPTIONS(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodOptions, path, meta, handler...)
}

func (w *Router) PATCH(path string, meta Meta, handler ...gin.HandlerFunc) RouterInfo {
	return w.Handle(http.MethodPatch, path, meta, handler...)
}

func (w *Router) Any(path string, meta Meta, handler ...gin.HandlerFunc) []RouterInfo {
	return w.Match([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
	}, path, meta, handler...)
}

func (w *Router) Match(method []string, path string, meta Meta, handler ...gin.HandlerFunc) []RouterInfo {
	var infos []RouterInfo
	for _, m := range method {
		infos = append(infos, w.Handle(m, path, meta, handler...))
	}
	return infos
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}
