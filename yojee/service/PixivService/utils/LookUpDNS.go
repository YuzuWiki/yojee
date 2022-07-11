package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const dnsQueryHost = "https://1.1.1.1/dns-query?name="

type dNsBody struct {
	Status int `json:"status"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
}

func LookUpDNS(host string) (err error, ip string) {
	if len(host) == 0 {
		return errors.New("invalid host"), ""
	}

	request, err := http.NewRequest(http.MethodGet, dnsQueryHost+host, nil)
	if err != nil {
		return
	} else {
		request.Header.Set("Accept", "application/dns-json")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("lookup dns fail, StatusCode is %d", response.StatusCode)), ""
	}

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	body := dNsBody{}
	if err := json.Unmarshal(data, &body); err != nil {
		return
	}

	for idx := range body.Answer {
		answer := body.Answer[idx]
		if answer.Type == 1 {
			return nil, answer.Data
		}
	}
	return errors.New("not found IP"), ""
}
