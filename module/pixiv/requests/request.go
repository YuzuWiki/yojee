package requests

import (
	"net/http"
	"strings"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

type requests struct {
	http.Client

	// Transport
	Transport pixiv.ITransport

	// 钩子函数..
	beforeHooks []pixiv.BeforeHook
	afterHooks  []pixiv.AfterHook
}

func doHooks[T *http.Request | *http.Response](hooks []func(T) error, body T) (err error) {
	for idx := range hooks {
		if err = hooks[idx](body); err != nil {
			return
		}
	}
	return
}

func newRequest(u, method string, query *pixiv.Query, params *pixiv.Params) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	u, err = pixiv.EncodeURL(u, query)
	if err != nil {
		return nil, err
	}

	body, err := pixiv.EncodeBody(params)
	if err == nil {
		req, err = http.NewRequest(method, u, body)
	} else {
		req, err = http.NewRequest(method, u, nil)
	}

	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *requests) do(method, u string, query *pixiv.Query, params *pixiv.Params) (resp *http.Response, err error) {
	var req *http.Request
	if req, err = newRequest(method, u, query, params); err != nil {
		return
	}

	if err = doHooks(r.beforeHooks, req); err != nil {
		return
	}

	switch strings.SplitN(strings.Replace(req.URL.Path, "/", "", 1), "/", 3)[0] {
	case "fanbox":
		resp, err = r.Transport.RoundTrip(req)
	default:
		resp, err = r.Client.Do(req)
	}

	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return nil, err
	}

	if err = doHooks(r.afterHooks, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *requests) Get(u string, query *pixiv.Query, params *pixiv.Params) (*http.Response, error) {
	return r.do(u, http.MethodGet, query, params)
}

func (r *requests) Post(u string, query *pixiv.Query, params *pixiv.Params) (*http.Response, error) {
	return r.do(u, http.MethodPost, query, params)
}

func (r *requests) Put(u string, query *pixiv.Query, params *pixiv.Params) (*http.Response, error) {
	return r.do(u, http.MethodPut, query, params)
}

func (r *requests) Delete(u string, query *pixiv.Query, params *pixiv.Params) (*http.Response, error) {
	return r.do(u, http.MethodDelete, query, params)
}

func (r *requests) SetProxy(proxyUrl string) error {
	return r.Transport.SetProxy(proxyUrl)
}

func (r *requests) UnSetProxy() error {
	return r.Transport.UnSetProxy()
}

func (r *requests) BeforeHooks(fns ...pixiv.BeforeHook) {
	for idx := range fns {
		r.beforeHooks = append(r.beforeHooks, fns[idx])
	}
}

func (r *requests) AfterHooks(fns ...pixiv.AfterHook) {
	for idx := range fns {
		r.afterHooks = append(r.afterHooks, fns[idx])
	}
}

func NewRequest() pixiv.IRequest {
	t := NewTransport()
	return &requests{
		Client: http.Client{
			Transport: t,
		},
		Transport: t,
	}
}
