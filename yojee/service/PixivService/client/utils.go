package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DNSBody struct {
	Status int `json:"Status"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
}

func lookupDNS(hostname string) (isOK bool, ip string) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://1.1.1.1/dns-query?name=%s", hostname), nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/dns-json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		ret := DNSBody{}
		if err := json.Unmarshal(data, &ret); err != nil {
			return
		}

		for i := range ret.Answer {
			ans := ret.Answer[i]
			if ans.Type == 1 {
				return true, ans.Data
			}
		}
	}
	return
}
