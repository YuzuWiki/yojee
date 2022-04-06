package client

import (
	"fmt"
	"net/http"
	"path"
)

func NewClient(phpSession string) *Client {
	c := new(Client)

	c.Host = Host
	c.SetPHPSESSID(phpSession)
	return c
}

type Client struct {
	http.Client

	// session host
	Host string
}

func (c *Client) endpointURL(subPath string) string {
	return fmt.Sprintf("https://%s", path.Join(c.Host, subPath))
}

func (c *Client) Get() {}

func (c *Client) Post() {}
