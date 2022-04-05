package client

import (
	"net/http"
	"path"
)

type Session struct {
	http.Client

	// session host
	Host string
}

func (s *Session) endpointURL(subPath string) string {
	return path.Join(s.Host, subPath)
}

func (s *Session) Get() {}

func (s *Session) Post() {}
