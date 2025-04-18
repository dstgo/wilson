// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.29.2
// source: configure/template/v1/configure_template_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTemplateCompareTemplate = "/wilson.api.configure.template.v1.Template/CompareTemplate"
const OperationTemplateCreateTemplate = "/wilson.api.configure.template.v1.Template/CreateTemplate"
const OperationTemplateCurrentTemplate = "/wilson.api.configure.template.v1.Template/CurrentTemplate"
const OperationTemplateGetTemplate = "/wilson.api.configure.template.v1.Template/GetTemplate"
const OperationTemplateListTemplate = "/wilson.api.configure.template.v1.Template/ListTemplate"
const OperationTemplatePreviewTemplate = "/wilson.api.configure.template.v1.Template/PreviewTemplate"
const OperationTemplateSwitchTemplate = "/wilson.api.configure.template.v1.Template/SwitchTemplate"

type TemplateHTTPServer interface {
	CompareTemplate(context.Context, *CompareTemplateRequest) (*CompareTemplateReply, error)
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateReply, error)
	CurrentTemplate(context.Context, *CurrentTemplateRequest) (*CurrentTemplateReply, error)
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateReply, error)
	ListTemplate(context.Context, *ListTemplateRequest) (*ListTemplateReply, error)
	PreviewTemplate(context.Context, *PreviewTemplateRequest) (*PreviewTemplateReply, error)
	SwitchTemplate(context.Context, *SwitchTemplateRequest) (*SwitchTemplateReply, error)
}

func RegisterTemplateHTTPServer(s *http.Server, srv TemplateHTTPServer) {
	r := s.Route("/")
	r.GET("/configure/api/v1/templates", _Template_ListTemplate0_HTTP_Handler(srv))
	r.GET("/configure/api/v1/template", _Template_GetTemplate0_HTTP_Handler(srv))
	r.GET("/configure/api/v1/template/current", _Template_CurrentTemplate0_HTTP_Handler(srv))
	r.POST("/configure/api/v1/template", _Template_CreateTemplate0_HTTP_Handler(srv))
	r.POST("/configure/api/v1/template/switch", _Template_SwitchTemplate0_HTTP_Handler(srv))
	r.POST("/configure/api/v1/template/compare", _Template_CompareTemplate0_HTTP_Handler(srv))
	r.POST("/configure/api/v1/template/preview", _Template_PreviewTemplate0_HTTP_Handler(srv))
}

func _Template_ListTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTemplateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateListTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTemplate(ctx, req.(*ListTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTemplateReply)
		return ctx.Result(200, reply)
	}
}

func _Template_GetTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTemplateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateGetTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTemplate(ctx, req.(*GetTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTemplateReply)
		return ctx.Result(200, reply)
	}
}

func _Template_CurrentTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrentTemplateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateCurrentTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CurrentTemplate(ctx, req.(*CurrentTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrentTemplateReply)
		return ctx.Result(200, reply)
	}
}

func _Template_CreateTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTemplateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateCreateTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateTemplate(ctx, req.(*CreateTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTemplateReply)
		return ctx.Result(200, reply)
	}
}

func _Template_SwitchTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SwitchTemplateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateSwitchTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.SwitchTemplate(ctx, req.(*SwitchTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SwitchTemplateReply)
		return ctx.Result(200, reply)
	}
}

func _Template_CompareTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompareTemplateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateCompareTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CompareTemplate(ctx, req.(*CompareTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompareTemplateReply)
		return ctx.Result(200, reply)
	}
}

func _Template_PreviewTemplate0_HTTP_Handler(srv TemplateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PreviewTemplateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplatePreviewTemplate)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.PreviewTemplate(ctx, req.(*PreviewTemplateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PreviewTemplateReply)
		return ctx.Result(200, reply)
	}
}

type TemplateHTTPClient interface {
	CompareTemplate(ctx context.Context, req *CompareTemplateRequest, opts ...http.CallOption) (rsp *CompareTemplateReply, err error)
	CreateTemplate(ctx context.Context, req *CreateTemplateRequest, opts ...http.CallOption) (rsp *CreateTemplateReply, err error)
	CurrentTemplate(ctx context.Context, req *CurrentTemplateRequest, opts ...http.CallOption) (rsp *CurrentTemplateReply, err error)
	GetTemplate(ctx context.Context, req *GetTemplateRequest, opts ...http.CallOption) (rsp *GetTemplateReply, err error)
	ListTemplate(ctx context.Context, req *ListTemplateRequest, opts ...http.CallOption) (rsp *ListTemplateReply, err error)
	PreviewTemplate(ctx context.Context, req *PreviewTemplateRequest, opts ...http.CallOption) (rsp *PreviewTemplateReply, err error)
	SwitchTemplate(ctx context.Context, req *SwitchTemplateRequest, opts ...http.CallOption) (rsp *SwitchTemplateReply, err error)
}

type TemplateHTTPClientImpl struct {
	cc *http.Client
}

func NewTemplateHTTPClient(client *http.Client) TemplateHTTPClient {
	return &TemplateHTTPClientImpl{client}
}

func (c *TemplateHTTPClientImpl) CompareTemplate(ctx context.Context, in *CompareTemplateRequest, opts ...http.CallOption) (*CompareTemplateReply, error) {
	var out CompareTemplateReply
	pattern := "/configure/api/v1/template/compare"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTemplateCompareTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateHTTPClientImpl) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...http.CallOption) (*CreateTemplateReply, error) {
	var out CreateTemplateReply
	pattern := "/configure/api/v1/template"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTemplateCreateTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateHTTPClientImpl) CurrentTemplate(ctx context.Context, in *CurrentTemplateRequest, opts ...http.CallOption) (*CurrentTemplateReply, error) {
	var out CurrentTemplateReply
	pattern := "/configure/api/v1/template/current"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTemplateCurrentTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateHTTPClientImpl) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...http.CallOption) (*GetTemplateReply, error) {
	var out GetTemplateReply
	pattern := "/configure/api/v1/template"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTemplateGetTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateHTTPClientImpl) ListTemplate(ctx context.Context, in *ListTemplateRequest, opts ...http.CallOption) (*ListTemplateReply, error) {
	var out ListTemplateReply
	pattern := "/configure/api/v1/templates"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTemplateListTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateHTTPClientImpl) PreviewTemplate(ctx context.Context, in *PreviewTemplateRequest, opts ...http.CallOption) (*PreviewTemplateReply, error) {
	var out PreviewTemplateReply
	pattern := "/configure/api/v1/template/preview"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTemplatePreviewTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateHTTPClientImpl) SwitchTemplate(ctx context.Context, in *SwitchTemplateRequest, opts ...http.CallOption) (*SwitchTemplateReply, error) {
	var out SwitchTemplateReply
	pattern := "/configure/api/v1/template/switch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTemplateSwitchTemplate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
