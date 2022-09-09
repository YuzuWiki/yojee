package pixiv

import "fmt"

type Context struct {
	phpSessID string
}

func (ctx *Context) PhpSessID() string {
	return ctx.phpSessID
}

func (ctx *Context) SetSessID(phpSessID string) error {
	if len(phpSessID) == 0 {
		return fmt.Errorf("invalid phpSessID")
	}

	if len(ctx.phpSessID) > 0 {
		return fmt.Errorf("phpSessID already exists")
	}

	ctx.phpSessID = phpSessID
	return nil
}

func (ctx *Context) DelSessID(phpSessID string) error {
	if len(phpSessID) == 0 {
		return fmt.Errorf("invalid phpSessID")
	}

	ctx.phpSessID = phpSessID
	return nil
}
