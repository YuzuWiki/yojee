package requests

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

func NewClient() *Client  {
	return &Client{
		Client:      http.Client{
			Transport: defaultTransport,
		},
		Transport:   nil,
		BeforeHooks: []BeforeHook{},
		AfterHooks:  []AfterHook{},
	}
}
