package client

import "net/http"

type beforeDo = func(req *http.Request) error

type afterDo = func(resp *http.Response) error

type Transport struct {
	before []beforeDo
	after  []afterDo

	transport http.RoundTripper
}

func (t *Transport) beforeRequest(req *http.Request) error {
	for _, fn := range t.before {
		if err := fn(req); err != nil {
			_ = req.Body.Close()
			return err
		}
	}
	return nil
}

func (t *Transport) afterRequest(resp *http.Response) error {
	for _, fn := range t.after {
		if err := fn(resp); err != nil {
			_ = resp.Body.Close()
			return err
		}
	}
	return nil
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	if err := t.beforeRequest(req); err != nil {
		return nil, err
	}

	resp, err = t.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if err := t.afterRequest(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
