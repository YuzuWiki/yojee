package apis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFollowingAPI_FindShow(t *testing.T) {
	var ctx = context.Background()

	f := FollowingAPI{}
	data, err := f.FindHide(ctx, 32835219, 24, 0)

	require.NoError(t, err)
	assert.NotEmpty(t, data.Users)
}

func TestFollowingAPI_FindHide(t *testing.T) {
	var ctx = context.Background()

	f := FollowingAPI{}
	data, err := f.FindHide(ctx, 32835219, 24, 0)

	require.NoError(t, err)
	assert.NotEmpty(t, data.Users)
}
