package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func (c Client) IsLoggedIn() (ret bool, err error) {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := c.Head(c.EndpointULR("/setting_user.php", nil).String())
	defer resp.Body.Close()
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, errors.New("pixiv client: unexpected  response for login test")
}

func (c *Client) ensureJar() {
	if c.Jar == nil {
		c.Jar, _ = cookiejar.New(nil)
	}
}

func (c *Client) Login(username string, password string) (err error) {
	c.ensureJar()

	resp, err := c.Get("https://accounts.pixiv.net/login?lang=zh")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	s := doc.Find(`input[name="post_key"]`)
	if len(s.Nodes) == 0 {
		return errors.New("pixiv client: can not found element for post key")
	}

	postKey, is_ok := s.Attr("value")
	if !is_ok {
		return errors.New("pixiv client: can not extract post key")
	}

	resp, err = c.PostForm("https://accounts.pixiv.net/api/login?lang=zh",
		url.Values{
			"pixiv_id": []string{username},
			"password": []string{password},
			"post_key": []string{postKey},
		})
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ParseAPIResult(resp.Body)
	if err != nil {
		return
	}

	if !body.Get("success").Exists() {
		return fmt.Errorf("pixiv client: login failed, error = %s", body.String())
	}
	return
}

func (c *Client) SetPHPSESSID(v string) {
	c.ensureJar()

	c.Jar.SetCookies(
		c.EndpointULR("", nil),
		[]*http.Cookie{
			{
				Domain: ".pixiv.net",
				Path:   "/",
				Name:   "PHPSESSID",
				Value:  v,
			},
		})
}
