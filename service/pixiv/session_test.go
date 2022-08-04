package pixiv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession_Session(t *testing.T) {
	se := NewSession()
	client := se.Get(os.Getenv(Phpsessid))

	proxyURL := os.Getenv("HTTP_PROXY")
	if len(proxyURL) > 0 {
		client.SetProxy(proxyURL)
	}

	resp, err := client.Get("https://www.pixiv.net", nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Header.Get("x-userid"))
}
