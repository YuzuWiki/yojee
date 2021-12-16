package following

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatchListAPI_Manga(t *testing.T) {
	api := WatchListAPI{}

	data, err := api.Manga(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Users)
	t.Log(data.TagTranslation)
	assert.NotEmpty(t, data.IllustSeries)
}

func TestWatchListAPI_Novel(t *testing.T) {
	api := WatchListAPI{}

	data, err := api.Novel(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Users)
	t.Log(data.TagTranslation)
	assert.NotEmpty(t, data.NovelSeries)
}
