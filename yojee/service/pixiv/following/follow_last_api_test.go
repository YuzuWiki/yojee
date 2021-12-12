package following

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastAPI_FollowLatestIllust(t *testing.T) {
	api := FollowLastAPI{}

	// test illust
	data, err := api.Illust(context.TODO(), ModAll, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Thumbnails.Illust)

	//  test  novel
	data, err = api.Novel(context.TODO(), ModAll, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Thumbnails.Novel)

}
