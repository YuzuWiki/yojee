package pixiv

import (
	_context "context"
	"errors"
	"strconv"
)

type ContextVar struct {
	_context.Context

	sessionID string
	uid       int64
}

func getUid(client RequestInterface) (int64, error) {
	resp, err := client.Get("https://"+PixivHost, nil, nil)
	if err != nil {
		return 0, err
	}

	uidStr := resp.Header.Get("x-userid")
	if len(uidStr) == 0 {
		return 0, errors.New("invalid login status")
	}

	return strconv.ParseInt(uidStr, 10, 64)
}

func (v *ContextVar) Client() RequestInterface {
	return Sessions.Get(v.sessionID)
}

func (v *ContextVar) PhpSessID() string {
	return v.sessionID
}

func (v *ContextVar) Uid() (int64, error) {
	if v.uid == 0 {

		uid, err := getUid(v.Client())
		if err != nil {
			return 0, err
		}

		v.uid = uid
	}
	return v.uid, nil
}
