package client

import (
	"github.com/like9th/yojee/yojee/service/pixiv/config"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
)

func lookupDNSOverHTTPS(dnsQueryURL string, hostname string) (ip string, err error) {
	req, err := http.NewRequest("GET", dnsQueryURL, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Set("name", hostname)
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/dns-json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var jsonData = gjson.ParseBytes(data)
	ip = jsonData.Get("Answer.#(type==1).data").String()
	return
}

var (
	Hosts              = map[string]string{}
	DNSQueryURL string = os.Getenv(config.PIXIV_DNS_QUERY_URL)
)

func init() {
	if DNSQueryURL == "" {
		DNSQueryURL = "https://1.1.1.1/dns-query"
	}
}

func resolveHostname(hostname string) (ip string, err error) {
	if v, isOk := Hosts[hostname]; isOk && v != "" {
		return v, nil
	}

	return lookupDNSOverHTTPS(DNSQueryURL, hostname)
}

func init() {
	Hosts["www.pixiv.net"], _ = resolveHostname("pixiv.net")
}
