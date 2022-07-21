package client

import (
	"net/http"
)

type (
	BeforeHook = func(req *http.Request) error
	AfterHook  = func(resp *http.Response) error
)

type Client struct {
	http.Client

	// Transport
	Transport *Transport

	// 钩子函数..
	BeforeHooks []BeforeHook
	AfterHooks  []AfterHook
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
