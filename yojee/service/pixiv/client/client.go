package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/like9th/yojee/yojee/service/pixiv/config"
	"github.com/tidwall/gjson"
)

type Client struct {
	ServerURL string
	http.Client
}

func (c Client) EndpointULR(path string, values *url.Values) *url.URL {
	s := c.ServerURL
	if s == "" {
		s = "https://www.pixiv.net"
	}

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	u.Path = path
	if values != nil {
		u.RawQuery = values.Encode()
	}
	return u
}

func (c *Client) GetWithContext(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func ParseAPIResult(r io.Reader) (ret gjson.Result, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	s := string(data)
	if !gjson.Valid(s) {
		err = fmt.Errorf("parse API error:　%s", s)
		return
	}

	ret = gjson.Parse(s)
	msg := ret.Get("message").String()
	if ret.Get("error").Bool() {
		err = fmt.Errorf("parse api error：%s", msg)
	}
	return
}

var (
	Default         = new(Client)
	DefaultUserAent = os.Getenv(config.PIXIV_USER_AGENT)
)

func init() {
	if DefaultUserAent == "" {
		DefaultUserAent = `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0`
	}

	if os.Getenv(config.PIXIV_BYPASS_SNI_BLOCKING) != "" {
		Default.BypassSNIBloccking()
	}

	Default.SetPHPSESSID(os.Getenv(config.PIXIV_PHPSESSID))
	Default.SetDefaultHeader("User-Agent", DefaultUserAent)
}
