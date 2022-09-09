package requests

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func (c *Client) ensureJar() {
	if c.Jar == nil {
		c.Jar, _ = cookiejar.New(nil)
	}
}

func (c *Client) SetCookies(rawURL string, cookies ...*http.Cookie) error {
	if len(rawURL) == 0 || len(cookies) == 0 {
		return fmt.Errorf("invalid params")
	}

	c.ensureJar()
	if !strings.HasPrefix(rawURL, "http") {
		rawURL = "https://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	c.Jar.SetCookies(u, cookies)
	return nil
}
