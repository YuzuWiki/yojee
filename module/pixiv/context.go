package pixiv

import (
	"errors"
	"strconv"
	"time"
)

type Context struct {
	sessionID string
	uid       int64

	values map[string]string
}

// getUid try return pixiv uid (if phpsessid is effective)
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

// Client return a pixiv client by phpsessid
func (ctx *Context) Client() RequestInterface {
	return Sessions.Get(ctx.sessionID)
}

// PhpSessID return phpsessid
func (ctx *Context) PhpSessID() string {
	return ctx.sessionID
}

// Uuid return pixiv uid (if phpsessid is effective)
func (ctx *Context) Uuid() (int64, error) {
	if ctx.uid == 0 {

		uid, err := getUid(ctx.Client())
		if err != nil {
			return 0, err
		}

		ctx.uid = uid
	}
	return ctx.uid, nil
}

// DeadLine always returns that there is no deadline (ok==false)
func (ctx *Context) DeadLine() (deadline time.Time, ok bool) {
	return
}

// Done always returns nil (chan which will wait forever),
func (ctx *Context) Done() <-chan struct{} {
	return nil
}

// Err always returns nil
func (ctx *Context) Err() error {
	return nil
}

// Value always return
func (ctx *Context) Value(key string) string {
	if value, isOk := ctx.values[key]; isOk {
		return value
	}
	return ""
}

func NewContext(sessid string) *Context {
	return &Context{
		sessionID: sessid,
	}
}
