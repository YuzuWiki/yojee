package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileAPI_All(t *testing.T) {
	api := ProfileAPI{}

	data, err := api.All(context.TODO(), 7038833)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Illusts)
}

func TestProfileAPI_Top(t *testing.T) {
	api := ProfileAPI{}

	// test: Manga + Illusts
	data, err := api.Top(context.TODO(), 7038833)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Illusts)

	// test: Novels
	data, err = api.Top(context.TODO(), 31171598)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Novels)
}
