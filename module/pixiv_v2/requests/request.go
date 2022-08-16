package requests

import "net/http"

type request struct {
	http.Client

	// Transport
	Transport *transport

	// 钩子函数..
	BeforeHooks []BeforeHook
	AfterHooks  []AfterHook
}

func (r *request) Get(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func (r *request) Post(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func (r *request) Put(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func (r *request) Delete(u string, query *Query, params *Params) (*http.Response, error) {
	return nil, nil
}

func NewRequest() RequestInterface {
	return &request{}
}
