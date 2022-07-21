package client

import (
	"bytes"
	"encoding/json"
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
		return nil, nil
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func (c *Client) do(u, method string, query *Query, params *Params) (*http.Response, error) {
	u, err := encodeURL(u, query)
	if err != nil {
		return nil, err
	}

	body, err := encodeBody(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	for idx := range c.beforeHooks {
		if err := c.beforeHooks[idx](req); err != nil {
			return nil, err
		}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	for idx := range c.AfterHooks {
		if err := c.AfterHooks[idx](resp); err != nil {
			_ = resp.Body.Close()
			return nil, err
		}
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
