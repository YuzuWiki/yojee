package main

import (
	"fmt"
	"net/http"
	"os"

	client "github.com/like9th/yojee/yojee/common/requests"
)

func getEnvAny(names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return ""
}

//
//func TestRequest() error {
//	//_ = os.Setenv("NO_PROXY", "www.pixiv.net")
//	fmt.Println("HTTP_PROXY: ", getEnvAny("HTTP_PROXY", "http_proxy"))
//	fmt.Println("HTTPS_PROXY: ", getEnvAny("HTTPS_PROXY", "https_proxy"))
//	fmt.Println("NO_PROXY: ", getEnvAny("NO_PROXY", "no_proxy"))
//
//	req, err := http.NewRequest(http.MethodGet, "https://www.pixiv.net/", nil)
//	if err != nil {
//		return err
//	}
//
//	transport := &client.Transport{}
//	transport.SetProxy(getEnvAny("HTTP_PROXY", "http_proxy"))
//
//	fmt.Println("begin: ", time.Now().UTC())
//	defer func() {
//		fmt.Println("end: ", time.Now().UTC())
//	}()
//	c := http.Client{
//		//Timeout:   5 * time.Second,
//		Transport: transport,
//	}
//
//	resp, err := c.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	//_, err := ioutil.ReadAll(resp.Body)
//	//if err != nil {
//	//	return err
//	//}
//
//	fmt.Println(resp.StatusCode)
//	return nil
//
//}
//
//func TestHost(host string) (*url.URL, error) {
//	u, err := url.Parse(host)
//	if err != nil {
//		return nil, err
//	}
//	return u, nil
//}
//
//func main() {
//	err := TestRequest()
//	if err != nil {
//		fmt.Println("TestRequest ERROR: ", err.Error())
//	} else {
//		fmt.Println("TestRequest SUCCESS")
//	}
//
//}

func main() {
	request := client.Client{
		AfterHooks: []client.AfterHook{
			func(resp *http.Response) error {
				fmt.Println("AfterHook: StatusCode =", resp.StatusCode)
				return nil
			},
		},
		BeforeHooks: []client.BeforeHook{
			func(req *http.Request) error {
				fmt.Println("BeforeHook: url = ", req.URL.String())
				return nil
			},
		},
	}
	request.SetProxy(getEnvAny("HTTP_PROXY", "http_proxy"))

	resp, err := request.Get("https://www.pixiv.net", nil, nil)
	if err == nil {
		fmt.Println("Success: ", resp.StatusCode)
		_ = resp.Body.Close()
	} else {
		fmt.Println("fail: ", err.Error())
	}
}
