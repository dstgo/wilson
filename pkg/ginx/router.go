package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maps"
	"net/http"
	"path"
	"slices"
)

var allowMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodDelete,
	http.MethodPut,
	http.MethodOptions,
	http.MethodHead,
}

type WalkRouteInfo struct {
	IsGroup bool
	Group   *WalkRouteInfo

	Method   string
	FullPath string
	Meta     Meta
}

type RouterHandler struct {
	chain gin.HandlersChain
	group *RouterGroup

	Method   string
	FullPath string
	Meta     Meta
}

func NewRouterGroup(group *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{
		root:     group,
		middles:  make(gin.HandlersChain, 0),
		groups:   make([]*RouterGroup, 0),
		handlers: make([]*RouterHandler, 0),
		Meta:     make(Meta),
		FullPath: group.BasePath(),
	}
}

type RouterGroup struct {
	// root group
	root *gin.RouterGroup
	// records the middleware chain
	middles gin.HandlersChain
	// subgroups
	groups []*RouterGroup
	// sub handlers
	handlers []*RouterHandler

	FullPath string
	Meta     Meta
}

// Attach handlers to the root group directly
func (r *RouterGroup) Attach(handlers ...gin.HandlerFunc) {
	r.root.Use(handlers...)
}

// Use append the handlers to chain, group's chain will be inherited by subgroup and sub handlers
func (r *RouterGroup) Use(handlers ...gin.HandlerFunc) {
	r.middles = append(r.middles, handlers...)
}

func (r *RouterGroup) Group(path string, subMeta Meta, handlers ...gin.HandlerFunc) *RouterGroup {
	subGroup := new(RouterGroup)
	if subMeta == nil {
		subMeta = make(Meta)
	}

	cloneGroupMeta := maps.Clone(r.Meta)

	// copy meta, overwrite group meta
	maps.Copy(cloneGroupMeta, subMeta)

	subRoot := r.root.Group(path, append(gin.HandlersChain{MetaHandler(subMeta)}, handlers...)...)

	subGroup.Meta = cloneGroupMeta
	subGroup.root = subRoot
	subGroup.FullPath = joinPaths(r.FullPath, path)
	subGroup.middles = append(subGroup.middles, r.middles...)

	r.groups = append(r.groups, subGroup)

	return subGroup
}

func (r *RouterGroup) Handle(method string, path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {

	subHandler := new(RouterHandler)

	if meta == nil {
		meta = make(Meta)
	}

	cloneGroupMeta := maps.Clone(r.Meta)

	// copy meta
	maps.Copy(cloneGroupMeta, meta)

	// inherit middlewares from group
	handlers = append(gin.HandlersChain{MetaHandler(meta)}, append(r.middles, handlers...)...)

	// handle route
	r.root.Handle(method, path, handlers...)

	subHandler.chain = handlers
	subHandler.group = r
	subHandler.Method = method
	subHandler.FullPath = joinPaths(r.FullPath, path)
	subHandler.Meta = cloneGroupMeta

	r.handlers = append(r.handlers, subHandler)

	return subHandler
}

func (r *RouterGroup) Match(methods []string, path string, meta Meta, handlers ...gin.HandlerFunc) []*RouterHandler {
	var hs []*RouterHandler
	for _, method := range methods {
		if !slices.Contains(allowMethods, method) {
			panic(fmt.Sprintf("not allowed method: %s", method))
		}
		hs = append(hs, r.Handle(method, path, meta, handlers...))
	}
	return hs
}

func (r *RouterGroup) GET(path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {
	return r.Handle(http.MethodGet, path, meta, handlers...)
}

func (r *RouterGroup) POST(path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {
	return r.Handle(http.MethodPost, path, meta, handlers...)
}

func (r *RouterGroup) DELETE(path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {
	return r.Handle(http.MethodDelete, path, meta, handlers...)
}

func (r *RouterGroup) PUT(path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {
	return r.Handle(http.MethodPut, path, meta, handlers...)
}

func (r *RouterGroup) OPTIONS(path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {
	return r.Handle(http.MethodOptions, path, meta, handlers...)
}

func (r *RouterGroup) HEAD(path string, meta Meta, handlers ...gin.HandlerFunc) *RouterHandler {
	return r.Handle(http.MethodHead, path, meta, handlers...)
}

func (r *RouterGroup) Any(path string, meta Meta, handlers ...gin.HandlerFunc) []*RouterHandler {
	return r.Match(allowMethods, path, meta, handlers...)
}

// Walk
// walk group and handlers info, include subgroup
func (r *RouterGroup) Walk(walk func(info WalkRouteInfo) error) error {
	infos := make([]WalkRouteInfo, 0, len(r.handlers)+1)

	// append group info
	groupInfo := WalkRouteInfo{
		IsGroup:  true,
		FullPath: r.FullPath,
		Meta:     r.Meta,
	}
	infos = append(infos, groupInfo)

	// append route info
	for _, handler := range r.handlers {
		infos = append(infos, WalkRouteInfo{
			IsGroup:  false,
			Method:   handler.Method,
			FullPath: handler.FullPath,
			Meta:     handler.Meta,
			Group:    &groupInfo,
		})
	}

	// walk infos
	for _, info := range infos {
		if err := walk(info); err != nil {
			return err
		}
	}

	// walk subgroup
	for _, group := range r.groups {
		if err := group.Walk(walk); err != nil {
			return err
		}
	}

	return nil
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
