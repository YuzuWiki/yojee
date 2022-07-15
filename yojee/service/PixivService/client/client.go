package client

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type (
	BeforeHook = func(req *http.Request) error
	AfterHook  = func(resp *http.Response) error

	HeaderOption struct {
		Key   string
		Value string
	}
)

type Client struct {
	http.Client

	// 钩子函数..
	beforeHooks []BeforeHook
	AfterHooks  []AfterHook
}

func (c *Client) SetHeader(options ...HeaderOption) {
	if len(options) == 0 {
		return
	}

	c.beforeHooks = append(
		c.beforeHooks,
		func(req *http.Request) error {
			for idx := range options {
				option := options[idx]

				req.Header.Set(option.Key, option.Value)
			}
			return nil
		},
	)
}

func (c *Client) ensureJar() {
	if c.Jar == nil {
		c.Jar, _ = cookiejar.New(nil)
	}
}

func (c *Client) SetCookies(rawURL string, cookies ...*http.Cookie) error {
	if len(rawURL) == 0 || len(cookies) == 0 {
		return errors.New("invalid params")
	}

	c.ensureJar()
	if !strings.HasPrefix(rawURL, "http") {
		rawURL = "https://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	c.Jar.SetCookies(u, cookies)
	return nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
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
