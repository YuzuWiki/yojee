package client

import "net/http"

type RequestOption = func(req *http.Request)

type RequestOptionsTransport struct {
	wrapped http.RoundTripper
	options []RequestOption
}

func (t *RequestOptionsTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	for _, option := range t.options {
		option(req)
	}

	if t.wrapped == nil {
		return http.DefaultTransport.RoundTrip(req)
	}

	return t.wrapped.RoundTrip(req)
}

func (c *Client) SetRequestOptions(options ...RequestOption) {
	c.Transport = &RequestOptionsTransport{
		wrapped: c.Transport,
		options: options,
	}
}

func (c *Client) SetDefaultHeader(key, value string) {
	c.SetRequestOptions(func(req *http.Request) {
		if len(req.Header[key]) > 0 {
			return
		}
		req.Header.Set(key, value)
	})
}
