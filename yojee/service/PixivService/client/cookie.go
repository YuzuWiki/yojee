package client

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func (c *Client) ensureJar() {
	if c.Jar == nil {
		c.Jar, _ = cookiejar.New(nil)
	}
}

func (c *Client) SetPHPSESSID(v string) {
	c.ensureJar()

	u, err := url.Parse(c.endpointURL(""))
	if err != nil {
		panic(err)
	}

	c.Jar.SetCookies(
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
