package following

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastAPI_FollowLatestIllust(t *testing.T) {
	api := LastAPI{}

	// test illust
	data, err := api.FollowLatestIllust(context.TODO(), Mold_Illust, ModAll, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Thumbnails.Illust)

	//  test  novel
	data, err = api.FollowLatestIllust(context.TODO(), Mold_Novel, ModAll, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Thumbnails.Novel)

}
