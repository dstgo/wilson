package http

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"

	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/pkg/strs"
)

type request struct {
	c        *config.Http
	request  *resty.Request
	logger   *log.Helper
	inputLog bool
}

type Request interface {
	DisableLog() Request
	Option(fn RequestFunc) Request
	Get(url string) (*response, error)
	Post(url string, data any) (*response, error)
	PostJson(url string, data any) (*response, error)
	Put(url string, data any) (*response, error)
	PutJson(url string, data any) (*response, error)
	Delete(url string) (*response, error)
	Do() (*response, error)
}

var httpClient *resty.Client

func Instance() *resty.Client {
	return httpClient
}

func Init(hc *config.Http, watcher config.Watcher) {

	if hc == nil {
		hc = &config.Http{
			EnableLog:        true,
			RetryCount:       3,
			RetryWaitTime:    100 * time.Millisecond,
			MaxRetryWaitTime: 3 * time.Second,
			Timeout:          10 * time.Second,
		}
	}

	client := resty.New()
	client.SetRetryWaitTime(hc.RetryWaitTime)
	client.SetRetryMaxWaitTime(hc.MaxRetryWaitTime)
	client.SetRetryCount(hc.RetryCount)
	client.SetTimeout(hc.Timeout)

	watcher("http", func(value config.Value) {
		err := value.Scan(hc)
		if err != nil {
			log.Errorf("watch http client config failed: %s", err.Error())
			return
		}
		log.Infof("watch http client config successfully")
	})

	httpClient = client
}

func NewRequest(conf *config.Http, logger *log.Helper) Request {
	req := httpClient.R()
	if conf.Server == "" {
		conf.Server = "kratosx http client"
	}
	req.Header.Set("User-Agent", conf.Server)
	return &request{
		c:        conf,
		request:  req,
		logger:   logger,
		inputLog: true,
	}
}

type RequestFunc func(*resty.Request)

func (h *request) DisableLog() Request {
	h.inputLog = false
	return h
}

func (h *request) Option(fn RequestFunc) Request {
	fn(h.request)
	return h
}

func (h *request) log(t int64, res *response) {
	if !(h.c.EnableLog && h.inputLog) {
		return
	}

	resData := res.Body()
	logs := []any{
		"type", "request",
		"method", h.request.Method,
		"url", h.request.URL,
		"header", h.request.Header,
		"body", h.request.Body,
		"cost", time.Now().UnixMilli() - t,
		"res", strs.BytesToString(resData),
	}
	if len(h.request.FormData) != 0 {
		logs = append(logs, "form-data", h.request.FormData)
	}
	if len(h.request.QueryParam) != 0 {
		logs = append(logs, "query", h.request.QueryParam)
	}
	h.logger.Infow(logs...)
}

func (h *request) Get(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Get(url)
	return res, res.err
}

func (h *request) Post(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Post(url)
	return res, res.err
}

func (h *request) PostJson(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Post(url)
	return res, res.err
}

func (h *request) Put(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.SetBody(data).Put(url)
	return res, res.err
}

func (h *request) PutJson(url string, data any) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.ForceContentType("application/json").SetBody(data).Put(url)
	return res, res.err
}

func (h *request) Delete(url string) (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Delete(url)
	return res, res.err
}

func (h *request) Do() (*response, error) {
	res := &response{}
	defer h.log(time.Now().UnixMilli(), res)
	res.response, res.err = h.request.Send()
	return res, res.err
}

type response struct {
	err      error
	response *resty.Response
}

func (r *response) Body() []byte {
	return r.response.Body()
}

func (r *response) Result(val any) error {
	return json.Unmarshal(r.response.Body(), val)
}
