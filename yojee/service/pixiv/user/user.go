package user

import (
	"context"
	"errors"
	"net/url"

	"github.com/like9th/yojee/yojee/service/pixiv/client"
	"github.com/like9th/yojee/yojee/service/pixiv/image"
)

type User struct {
	ID     string
	Name   string
	Avatar image.URLs
}

func (u *User) Fetch(ctx context.Context) (err error) {
	if u.ID == "" {
		return errors.New("pixiv user: miss id")
	}

	clit := client.For(ctx)
	resp, err := clit.GetWithContext(ctx, clit.EndpointULR("/ajax/user/"+u.ID, nil).String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := client.ParseAPIResult(resp.Body)
	if err != nil {
		return
	}

	u.Name = body.Get("name").String()
	u.Avatar.Mini = body.Get("image").String()
	u.Avatar.Thumb = body.Get("imageBig").String()
	return
}

func (u User) URL(ctx context.Context) *url.URL {
	return client.For(ctx).EndpointULR("/user/"+u.ID, nil)
}
