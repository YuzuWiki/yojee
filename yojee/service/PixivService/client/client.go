package client

import (
	"net/http"
)

type Request struct {
	http.Client
}

func (r *Request) SetHeard() {
}

func (r *Request) SetCookies() {
}

func (r *Request) initTransport() {
}
