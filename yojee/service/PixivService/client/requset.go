package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type (
	Query  = url.Values          // request data
	Params = map[string]struct{} // request body
)

// encode url
func encodeURL(u string, data *Query) (string, error) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if data != nil {
		URL.RawQuery = data.Encode()
	}

	return URL.String(), nil
}

func encodeBody(params *Params) (*bytes.Buffer, error) {
	if params == nil {
		return nil, errors.New("params is nil")
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func newRequest(u, method string, query *Query, params *Params) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	u, err = encodeURL(u, query)
	if err != nil {
		return nil, err
	}

	body, err := encodeBody(params)
	if err == nil {
		req, err = http.NewRequest(method, u, body)
	} else {
		req, err = http.NewRequest(method, u, nil)
	}

	if err != nil {
		return nil, err
	} else {
		return req, nil
	}
}

func doHooks[T *http.Request | *http.Response](hooks []func(T) error, body T) error {
	if hooks != nil {
		for idx := range hooks {
			if err := hooks[idx](body); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) do(method, u string, query *Query, params *Params) (*http.Response, error) {
	req, err := newRequest(method, u, query, params)
	if err != nil {
		return nil, err
	}

	if err := doHooks(c.BeforeHooks, req); err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := doHooks(c.AfterHooks, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Get(u string, query *Query, params *Params) (*http.Response, error) {
	return c.do(u, http.MethodGet, query, params)
}

func (c *Client) Post(u string, query *Query, params *Params) (*http.Response, error) {
	return c.do(u, http.MethodPost, query, params)
}

func (c *Client) Put(u string, query *Query, params *Params) (*http.Response, error) {
	return c.do(u, http.MethodPut, query, params)
}

func (c *Client) Delete(u string, query *Query, params *Params) (*http.Response, error) {
	return c.do(u, http.MethodDelete, query, params)
}
