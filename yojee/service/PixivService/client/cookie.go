package client

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func (r *Requests) ensureJar() {
	if r.Jar == nil {
		r.Jar, _ = cookiejar.New(nil)
	}
}

func (r *Requests) SetPHPSESSID(v string) {
	r.ensureJar()

	u, err := url.Parse(r.endpointURL(""))
	if err != nil {
		panic(err)
	}

	r.Jar.SetCookies(
		u,
		[]*http.Cookie{
			{
				Domain: Domain,
				Path:   "/",
				Name:   "PHPSESSID",
				Value:  v,
			},
		})
}
