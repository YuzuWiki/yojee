package client

import (
	"net/http"
	"sync"
)

var (
	// byPassSNI
	byPassSNITransport *http.Transport
)

type BeforeHook = func(req *http.Request) error

type AfterHook = func(resp *http.Response) error

type Transport struct {
	http.Transport

	// hook
	beforeHooks []BeforeHook
	afterHooks  []AfterHook

	// mutex
	mu sync.Mutex
}

func (t *Transport) beforeRequest(req *http.Request) error {
	for _, before := range t.beforeHooks {
		if err := before(req); err != nil {
			_ = req.Body.Close()
			return err
		}
	}
	return nil
}

func (t *Transport) afterRequest(resp *http.Response) error {
	for _, after := range t.afterHooks {
		if err := after(resp); err != nil {
			_ = resp.Body.Close()
			return err
		}
	}
	return nil
}

func (t *Transport) byPassSNI(req *http.Request) (resp *http.Response, err error) {
	if byPassSNITransport == nil {
		t.mu.Lock()
		if byPassSNITransport == nil {
			byPassSNITransport = &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				//DialTLSContext: DialTLSContext,  #  TODO: remote error: tls: handshake failure
			}
		}
		t.mu.Unlock()
	}
	return byPassSNITransport.RoundTrip(req)
}

func (t *Transport) roundTrip(req *http.Request) (resp *http.Response, err error) {
	switch req.Host {
	case PIXIV_HOST, PXIMG_HOST:
		return t.byPassSNI(req)
	default:
		// return http.DefaultTransport.RoundTrip(req)
		return t.Transport.RoundTrip(req)
	}
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	if err := t.beforeRequest(req); err != nil {
		return nil, err
	}

	resp, err = t.roundTrip(req)
	if err != nil {
		return nil, err
	}

	if err := t.afterRequest(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
