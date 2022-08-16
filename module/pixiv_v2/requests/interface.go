package requests

import (
	"net/http"
	"net/url"
)

type (
	Query        = url.Values          // request data
	Params       = map[string]struct{} // request body
	BeforeHook   = func(req *http.Request) error
	AfterHook    = func(resp *http.Response) error
	HeaderOption struct {
		Key   string
		Value string
	}
)

type TransportInterface interface {
	SetProxy(string) error
	UnSetProxy() error
}

type RequestInterface interface {
	Get(string, *Query, *Params) (*http.Response, error)
	Post(string, *Query, *Params) (*http.Response, error)
	Put(string, *Query, *Params) (*http.Response, error)
	Delete(string, *Query, *Params) (*http.Response, error)
}
