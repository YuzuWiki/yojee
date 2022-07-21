package requests

import "net/http"

type (
	HeaderOption struct {
		Key   string
		Value string
	}
)

func (c *Client) SetHeader(options ...HeaderOption) {
	if len(options) == 0 {
		return
	}

	c.BeforeHooks = append(
		c.BeforeHooks,
		func(req *http.Request) error {
			for idx := range options {
				option := options[idx]

				req.Header.Set(option.Key, option.Value)
			}
			return nil
		},
	)
}
