package main

import (
	"github.com/like9th/yojee/yojee/common/requests"
	"net/http"
	"sync"
)

var Sessions *Session

type ClientInterface interface {
	SetProxy(string)
	UnSetProxy()

	Get(string, *requests.Query, *requests.Params) (*http.Response, error)
	Post(string, *requests.Query, *requests.Params) (*http.Response, error)
	Put(string, *requests.Query, *requests.Params) (*http.Response, error)
	Delete(string, *requests.Query, *requests.Params) (*http.Response, error)
}

func newClient(phpSessid string) *requests.Client {
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
			Value:  phpSessid,
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

func (s *Session) Get(phpSessid string) ClientInterface {
	client, isOk := s.pool[phpSessid]
	if !isOk {
		s.m.Lock()
		if _, isOk := s.pool[phpSessid]; !isOk {
			client = newClient(phpSessid)
			s.pool[phpSessid] = client
		}
		s.m.Unlock()
	}
	return client
}

func (s *Session) Delete(phpSessid string) {
	delete(s.pool, phpSessid)
}

func NewSession() *Session {
	return &Session{pool: map[string]ClientInterface{}}
}

func init() {
	if Sessions == nil {
		Sessions = NewSession()
	}
}
