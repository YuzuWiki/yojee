package apis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookmarkAPI_FindShow(t *testing.T) {
	var ctx = context.Background()

	api := BookmarkAPI{}
	data, err := api.FindShow(ctx, 32835219, "アークナイツ", 0, 48)

	require.NoError(t, err)
	assert.NotEmpty(t, data.Works)
}

func TestBookmarkAPI_FindHide(t *testing.T) {
	var ctx = context.Background()

	api := BookmarkAPI{}
	data, err := api.FindHide(ctx, 32835219, "", 0, 48)

	require.NoError(t, err)
	assert.NotEmpty(t, data.Works)
}
