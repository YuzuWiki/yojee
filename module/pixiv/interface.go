package pixiv

import (
	"net/http"
	"net/url"
)

type (
	Query        = url.Values          // requests data
	Params       = map[string]struct{} // requests body
	BeforeHook   = func(req *http.Request) error
	AfterHook    = func(resp *http.Response) error
	HeaderOption struct {
		Key   string
		Value string
	}
)

type ITransport interface {
	SetProxy(string) error
	UnSetProxy() error
}

type IClient interface {
	Get(string, *Query, *Params) (*http.Response, error)
	Post(string, *Query, *Params) (*http.Response, error)
	Put(string, *Query, *Params) (*http.Response, error)
	Delete(string, *Query, *Params) (*http.Response, error)
}

type IRequest interface {
	IClient

	SetCookies(string, ...*http.Cookie) error
	SetHeader(...HeaderOption)
	SetProxy(string) error
	UnSetProxy() error

	BeforeHooks(...BeforeHook)
	AfterHooks(...AfterHook)
}

type IContext interface {
	PhpSessID() string
}

type ISession interface {
	Default() (IClient, error)
	New(string) (IClient, error)
	Remove(string)
}
