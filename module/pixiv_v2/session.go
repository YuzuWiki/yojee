package pixiv_v2

import (
	"github.com/YuzuWiki/yojee/module/pixiv_v2/requests"
	"net/http"
	"os"
	"strings"
	"sync"
)

var Session ISession

type IClient interface {
	Get(string, *requests.Query, *requests.Params) (*http.Response, error)
	Post(string, *requests.Query, *requests.Params) (*http.Response, error)
	Put(string, *requests.Query, *requests.Params) (*http.Response, error)
	Delete(string, *requests.Query, *requests.Params) (*http.Response, error)
}

func newClient(sessionIds ...string) (IClient, error) {
	c := requests.NewRequest()

	// set default header
	c.SetHeader(
		requests.HeaderOption{Key: "User-Agent", Value: UserAgent},
		requests.HeaderOption{Key: "referer", Value: "https://" + PixivHost},
	)

	// set cookie
	if len(sessionIds) > 0 {
		err := c.SetCookies(
			PixivHost,
			&http.Cookie{
				Name:   Phpsessid,
				Value:  sessionIds[0],
				Path:   "/",
				Domain: PixivDomain,
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

type ISession interface {
	Default() (IClient, error)
	New(string) (IClient, error)
	Remove(string)
}

type session struct {
	pool map[string]IClient

	m sync.Mutex
}

func (s *session) new(sessionId string) (IClient, error) {
	sessionId = strings.TrimSpace(sessionId)
	if c, isOk := s.pool[sessionId]; isOk {
		return c, nil
	}

	s.m.Lock()
	defer s.m.Unlock()

	if c, isOk := s.pool[sessionId]; isOk {
		return c, nil
	}

	c, err := newClient(sessionId)
	if err != nil {
		return nil, err
	}

	s.pool[sessionId] = c
	return c, nil
}

func (s *session) Default() (IClient, error) {
	return s.new("_DEFAULT")
}

func (s *session) New(sessionId string) (IClient, error) {
	return s.new(sessionId)
}

func (s *session) Remove(sessionId string) {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.pool, sessionId)
}

func init() {
	if Session == nil {
		Session = &session{}
	}
}