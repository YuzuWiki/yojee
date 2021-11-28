package artwork

import (
	"context"
	"testing"
)

func TestFollowingAPI_FindShow(t *testing.T) {
	var ctx = context.Background()

	f := FollowingAPI{}
	data, err := f.FindShow(ctx, 32835219, 24,0)
	t.Log(err)
	t.Log(data)
}
