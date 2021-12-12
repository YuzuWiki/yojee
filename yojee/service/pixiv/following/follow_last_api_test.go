package following

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFollowLastAPI_Illust(t *testing.T) {
	api := FollowLastAPI{}

	data, err := api.Illust(context.TODO(), ModAll, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Thumbnails.Illust)
}

func TestFollowLastAPI_Novel(t *testing.T) {
	api := FollowLastAPI{}

	data, err := api.Novel(context.TODO(), ModAll, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Thumbnails.Novel)
}
