package requests

import (
	"net/http"

	"github.com/YuzuWiki/yojee/module/pixiv_v2"
)

type requests struct {
	http.Client

	// Transport
	Transport *transport

	// 钩子函数..
	beforeHooks []pixiv_v2.BeforeHook
	afterHooks  []pixiv_v2.AfterHook
}

func doHooks[T *http.Request | *http.Response](hooks []func(T) error, body T) (err error) {
	for idx := range hooks {
		if err = hooks[idx](body); err != nil {
			return
		}
	}
	return
}

func newRequest(u, method string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	u, err = pixiv_v2.EncodeURL(u, query)
	if err != nil {
		return nil, err
	}

	body, err := pixiv_v2.EncodeBody(params)
	if err == nil {
		req, err = http.NewRequest(method, u, body)
	} else {
		req, err = http.NewRequest(method, u, nil)
	}

	if err != nil {
		return nil, err
	} else {
		return req, nil
	}
}

func (r *requests) do(method, u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (resp *http.Response, err error) {
	var req *http.Request
	if req, err = newRequest(method, u, query, params); err != nil {
		return
	}

	if err = doHooks(r.beforeHooks, req); err != nil {
		return
	}

	if resp, err = r.Client.Do(req); err != nil {
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

func (r *requests) Get(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return r.do(u, http.MethodGet, query, params)
}

func (r *requests) Post(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return r.do(u, http.MethodPost, query, params)
}

func (r *requests) Put(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return r.do(u, http.MethodPut, query, params)
}

func (r *requests) Delete(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return r.do(u, http.MethodDelete, query, params)
}

func (r *requests) SetProxy(proxyUrl string) error {
	return r.Transport.SetProxy(proxyUrl)
}

func (r *requests) UnSetProxy() error {
	return r.Transport.UnSetProxy()
}

func (r *requests) BeforeHooks(fns ...pixiv_v2.BeforeHook) {
	for idx := range fns {
		r.beforeHooks = append(r.beforeHooks, fns[idx])
	}
}

func (r *requests) AfterHooks(fns ...pixiv_v2.AfterHook) {
	for idx := range fns {
		r.afterHooks = append(r.afterHooks, fns[idx])
	}
}

func NewRequest() pixiv_v2.IRequest {
	return &requests{}
}
