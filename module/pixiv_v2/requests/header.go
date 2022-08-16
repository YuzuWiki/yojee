package requests

import "net/http"

func (r *request) SetHeader(options ...HeaderOption) {
	if len(options) == 0 {
		return
	}

	r.BeforeHooks = append(
		r.BeforeHooks,
		func(req *http.Request) error {
			for _, option := range options {
				req.Header.Set(option.Key, option.Value)
			}
			return nil
		},
	)
}
