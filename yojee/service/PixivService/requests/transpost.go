package requests

import (
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/http/httpproxy"
)

var (
	defaultTransport http.RoundTripper
)

type Transport struct {
	// TODO: 路由分发
	http.Transport

	// 代理设置
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

func (t *Transport) SetProxy(proxyUrl string) {
	if len(proxyUrl) == 0 || t.ProxyURl == proxyUrl {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	t.ProxyURl = proxyUrl

	if t.Proxy == nil {
		t.Transport.Proxy = t.proxyFromUrl
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
		return defaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	return t.roundTrip(req)
}

func NewTransport() *Transport {
	return &Transport{
		ProxyURl: "",
		mu:       sync.Mutex{},

		Transport: http.Transport{
			DisableKeepAlives: true,
		},
	}
}

func (c *Client) SetProxy(proxy string) {
	if c.Transport == nil {
		c.Transport = NewTransport()
		c.Client.Transport = c.Transport
	}

	c.Transport.SetProxy(proxy)
}

func (c *Client) UnSetProxy() {
	if c.Transport != nil {
		c.Transport.UnSetProxy()
	}
}

func init() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	defaultTransport = &http.Transport{
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
