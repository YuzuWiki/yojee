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

	// 钩子函数..
	BeforeHooks []BeforeHook
	AfterHooks  []AfterHook
}
