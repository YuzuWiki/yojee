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
	BeforeHooks []pixiv_v2.BeforeHook
	AfterHooks  []pixiv_v2.AfterHook
}

func (r *requests) Get(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) Post(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) Put(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) Delete(u string, query *pixiv_v2.Query, params *pixiv_v2.Params) (*http.Response, error) {
	return nil, nil
}

func (r *requests) SetProxy(proxyUrl string) error {
	return r.Transport.SetProxy(proxyUrl)
}

func (r *requests) UnSetProxy() error {
	return r.Transport.UnSetProxy()
}

func NewRequest() pixiv_v2.IRequest {
	return &requests{}
}
