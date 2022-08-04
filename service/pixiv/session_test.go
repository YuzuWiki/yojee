package pixiv

import (
	"fmt"
	"os"
	"testing"
)

func TestSession_Session(t *testing.T) {
	//se := NewSession()
	fmt.Println("*********: ", os.Getenv(Phpsessid))
	//client := se.Get(os.Getenv(Phpsessid))
	//
	//proxyURL := os.Getenv("HTTP_PROXY")
	//if len(proxyURL) > 0 {
	//	client.SetProxy("socks5://127.0.0.1:27001")
	//}
	//
	//resp, err := client.Get("https://www.pixiv.net", nil, nil)
	//assert.NoError(t, err)
	//fmt.Println(os.Getenv(Phpsessid))
	//assert.NotEmpty(t, resp.Header.Get("x-userid"))
}
