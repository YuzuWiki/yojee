package client

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Request struct {
	http.Client
}

func (r *Request) SetHeard() {
}

func (r *Request) ensureJar() {
	if r.Jar == nil {
		r.Jar, _ = cookiejar.New(nil)
	}
}

func (r *Request) SetCookies(rawURL string, cookies ...*http.Cookie) error {
	if len(rawURL) == 0 || len(cookies) == 0 {
		return errors.New("invalid params")
	}

	r.ensureJar()
	if !strings.HasPrefix(rawURL, "http") {
		rawURL = "https://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	r.Jar.SetCookies(u, cookies)
	return nil
}

func (r *Request) initTransport() {
}
