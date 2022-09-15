package requests

import (
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

func newClient(sessionIds ...string) (pixiv.IClient, error) {
	c := NewRequest()

	// set default header
	c.SetHeader(
		pixiv.HeaderOption{Key: "User-Agent", Value: pixiv.UserAgent},
		pixiv.HeaderOption{Key: "referer", Value: "https://" + pixiv.PixivHost},
	)

	// set cookie
	if len(sessionIds) > 0 {
		err := c.SetCookies(
			pixiv.PixivHost,
			&http.Cookie{
				Name:   pixiv.Phpsessid,
				Value:  sessionIds[0],
				Path:   "/",
				Domain: pixiv.PixivDomain,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	// set proxy
	if proxyUrl := os.Getenv("Proxy_URL"); len(proxyUrl) > 0 {
		if err := c.SetProxy(proxyUrl); err != nil {
			return nil, err
		}
	}
	return c, nil
}

type session struct {
	pool map[string]pixiv.IClient

	m sync.Mutex
}

func (s *session) new(sessionId string) (c pixiv.IClient, err error) {
	sessionId = strings.TrimSpace(sessionId)
	if c, isOk := s.pool[sessionId]; isOk {
		return c, nil
	}

	s.m.Lock()
	defer s.m.Unlock()

	if c, isOk := s.pool[sessionId]; isOk {
		return c, nil
	}

	if c, err = newClient(sessionId); err != nil {
		return
	}

	s.pool[sessionId] = c
	return
}

func (s *session) Default() (pixiv.IClient, error) {
	return s.new("_DEFAULT")
}

func (s *session) New(sessionId string) (pixiv.IClient, error) {
	return s.new(sessionId)
}

func (s *session) Remove(sessionId string) {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.pool, sessionId)
}

func NewSession() pixiv.ISession {
	return &session{pool: map[string]pixiv.IClient{}}
}
