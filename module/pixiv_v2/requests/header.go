package requests

import (
	"net/http"

	"github.com/YuzuWiki/yojee/module/pixiv_v2"
)

func (r *requests) SetHeader(options ...pixiv_v2.HeaderOption) {
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
