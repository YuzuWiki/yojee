package main

import (
	"net/http"

	"github.com/like9th/yojee/yojee/common/requests"
)

var Sessions *Session

type ClientInterface interface {
	SetProxy(string)
	UnSetProxy(string)

	Get(string, *requests.Query, *requests.Params) (*http.Response, error)
	Post(string, *requests.Query, *requests.Params) (*http.Response, error)
	Put(string, *requests.Query, *requests.Params) (*http.Response, error)
	Delete(string, *requests.Query, *requests.Params) (*http.Response, error)
}

type Session struct {
	pool map[string]*ClientInterface
}

func (s *Session) Get(phpSessid string) *ClientInterface {
	client, isOk := s.pool[phpSessid]
	if !isOk {
		// TODO: new client
	}
	return client
}
//
//func newClient() {
//	PixivClient := &requests.Client{
//		Client:      http.Client{},
//		Transport:   nil,
//		BeforeHooks: []requests.BeforeHook{},
//		AfterHooks:  []requests.AfterHook{},
//	}
//
//	PixivClient.SetHeader(
//		requests.HeaderOption{
//			Key:   "User-Agent",
//			Value: UserAgent,
//		})
//}

func NewSession() *Session  {
	return &Session{pool: map[string]*ClientInterface{}}
}

func init(){
	if Sessions == nil {
		Sessions = NewSession()
	}
}