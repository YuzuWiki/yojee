package pixiv

import (
	"net/http"
	"os"
	"sync"

	"github.com/YuzuWiki/yojee/common/requests"
)

var Sessions *Session

type _RequestMethod func(string, *requests.Query, *requests.Params) (*http.Response, error)

type RequestInterface interface {
	Get(string, *requests.Query, *requests.Params) (*http.Response, error)
	Post(string, *requests.Query, *requests.Params) (*http.Response, error)
	Put(string, *requests.Query, *requests.Params) (*http.Response, error)
	Delete(string, *requests.Query, *requests.Params) (*http.Response, error)
}

type ClientInterface interface {
	RequestInterface

	SetProxy(string)
	UnSetProxy()
}

func newClient(sessionID string) *requests.Client {
	// new client
	c := requests.Client{
		Client:      http.Client{},
		Transport:   nil,
		BeforeHooks: []requests.BeforeHook{},
		AfterHooks:  []requests.AfterHook{},
	}

	// set default header
	c.SetHeader(
		requests.HeaderOption{
			Key:   "User-Agent",
			Value: UserAgent,
		},
		requests.HeaderOption{
			Key:   "referer",
			Value: "https://" + PixivHost,
		})

	// set cookie
	if err := c.SetCookies(
		PixivHost,
		&http.Cookie{
			Name:   Phpsessid,
			Value:  sessionID,
			Path:   "/",
			Domain: PixivDomain,
		},
	); err != nil {
		panic(err.Error())
	}

	if proxyURL := os.Getenv("PROXY_URL"); len(proxyURL) > 0 {
		c.SetProxy(proxyURL)
	}
	return &c
}

type Session struct {
	pool map[string]RequestInterface

	m sync.Mutex
}

func (s *Session) Get(sessionID string) RequestInterface {
	if client, isOk := s.pool[sessionID]; isOk {
		return client
	}

	s.m.Lock()
	defer s.m.Unlock()
	if client, isOk := s.pool[sessionID]; isOk {
		return client
	}

	client := newClient(sessionID)
	s.pool[sessionID] = client
	return client
}

func (s *Session) Delete(sessionID string) {
	delete(s.pool, sessionID)
}

func NewSession() *Session {
	return &Session{pool: map[string]RequestInterface{}}
}

func init() {
	if Sessions == nil {
		Sessions = NewSession()
	}
}
