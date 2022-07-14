package client

import (
	"golang.org/x/net/http/httpproxy"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	DefaultTransport http.RoundTripper
)

type Transport struct {
	http.Transport

	// 代理设置
	proxy    func(*http.Request) (*url.URL, error)
	ProxyURl string

	// mutex
	mu sync.Mutex
}

func (t *Transport) proxyFromUrl(req *http.Request) (*url.URL, error) {
	cnf := &httpproxy.Config{
		HTTPProxy:  t.ProxyURl,
		HTTPSProxy: t.ProxyURl,
		NoProxy:    "",
		CGI:        false,
	}
	return cnf.ProxyFunc()(req.URL)
}

func (t *Transport) SetProxy(url string) {
	if len(url) == 0 || t.ProxyURl == url {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	t.ProxyURl = url

	if t.Proxy == nil {
		t.Proxy = t.proxyFromUrl
	}
}

func (t *Transport) UnSetProxy() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ProxyURl = ""
	t.Proxy = nil
}

func (t *Transport) roundTrip(req *http.Request) (resp *http.Response, err error) {
	if len(t.ProxyURl) == 0 {
		return DefaultTransport.RoundTrip(req)
	}
	return t.RoundTrip(req)
}

func NewTransport() *Transport {
	return &Transport{
		proxy:    nil,
		ProxyURl: "",
	}
}

func init() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	DefaultTransport = &http.Transport{
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
