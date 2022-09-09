package requests

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/http/httpproxy"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

var _defaultTransport http.RoundTripper

type transport struct {
	http.Transport

	// mutex
	mu sync.Mutex
}

func proxyFromUrl(proxyUrl string) func(*http.Request) (*url.URL, error) {
	return func(req *http.Request) (*url.URL, error) {
		cnf := &httpproxy.Config{
			HTTPProxy:  proxyUrl,
			HTTPSProxy: proxyUrl,
			NoProxy:    "",
			CGI:        false,
		}
		return cnf.ProxyFunc()(req.URL)
	}
}

func (t *transport) SetProxy(proxyUrl string) error {
	if len(proxyUrl) == 0 {
		return fmt.Errorf("invalid proxy url")
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Proxy == nil {
		t.Transport.Proxy = proxyFromUrl(proxyUrl)
	}
	return nil
}

func (t *transport) UnSetProxy() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Proxy = nil
	return nil
}

func (t *transport) roundTrip(req *http.Request) (resp *http.Response, err error) {
	if t.Proxy == nil {
		return _defaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	return t.roundTrip(req)
}

func NewTransport() pixiv.ITransport {
	return &transport{
		mu: sync.Mutex{},

		Transport: http.Transport{
			DisableKeepAlives: true,
		},
	}
}

func init() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	_defaultTransport = &http.Transport{
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
