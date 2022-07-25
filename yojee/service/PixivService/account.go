package main

import (
	"errors"
	"strconv"
)

type Account struct {
	phpSessid string

	uid    int64
	client ClientInterface
}

func (a Account) getUid() (int64, error)  {
	resp, err := a.client.Get("https://"+PixivHost, nil, nil)
	if err != nil {
		return 0, err
	}

	uid := resp.Header.Get("x-userid")
	if len(uid) == 0 {
		return 0, errors.New("invalid request")
	}

	return strconv.ParseInt(uid, 10, 64)
}


func (a *Account) Uid() (int64, error) {
	if a.uid == 0 {
		uid, err := a.getUid()
		if err != nil {
			return 0, err
		}

		a.uid = uid
	}
	return a.uid, nil
}



