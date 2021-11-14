package client

import "context"

type contextKey struct {}


func For(ctx context.Context) *Client {
	v, _ := ctx.Value(contextKey{}).(*Client)
	if v == nil {
		return Default
	}
	return v
}

func With(ctx context.Context, v *Client) context.Context {
	return context.WithValue(ctx, contextKey{}, v)
}