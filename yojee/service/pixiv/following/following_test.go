package following

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFollowAPI_FindShow(t *testing.T) {
	api := FollowAPI{}

	data, err := api.FindHide(context.TODO(), 32835219, 24, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, data.Users)
}

func TestFollowAPI_FindHide(t *testing.T) {
	api := FollowAPI{}

	data, err := api.FindHide(context.TODO(), 32835219, 24, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, data.Users)
}
