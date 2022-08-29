package requests

import "net/http"

type requests struct {
	http.Client

	// Transport
	Transport *transport

	// 钩子函数..
	BeforeHooks []BeforeHook
	AfterHooks  []AfterHook
}

func (r *requests) Get(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) Post(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) Put(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) Delete(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func NewRequest() IRequest {
	return &requests{}
}
