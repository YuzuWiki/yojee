package pixivService

import (
	"github.com/like9th/yojee/yojee/common/requests"
	"net/http"
	"sync"
)

var Sessions *Session

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
	return &c
}

type Session struct {
	pool map[string]ClientInterface

	m sync.Mutex
}

func (s *Session) Get(sessionID string) RequestInterface {
	client, isOk := s.pool[sessionID]
	if !isOk {
		s.m.Lock()
		if _, isOk := s.pool[sessionID]; !isOk {
			client = newClient(sessionID)
			s.pool[sessionID] = client
		}
		s.m.Unlock()
	}
	return client
}

func (s *Session) Delete(sessionID string) {
	delete(s.pool, sessionID)
}

func NewSession() *Session {
	return &Session{pool: map[string]ClientInterface{}}
}

func init() {
	if Sessions == nil {
		Sessions = NewSession()
	}
}
