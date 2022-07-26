package pixivService

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSession_Session(t *testing.T) {
	se := NewSession()
	client := se.Get(os.Getenv(Phpsessid))

	_ = os.Setenv("HTTP_PROXY", "socks://127.0.0.1:27001")
	proxyURL := os.Getenv("HTTP_PROXY")
	if len(proxyURL) > 0 {
		client.SetProxy("socks5://127.0.0.1:27001")

	}

	resp, err := client.Get("https://www.pixiv.net", nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Header.Get("x-userid"))
}
