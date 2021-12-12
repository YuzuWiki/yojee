package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookMarkAPI_FindShow(t *testing.T) {
	api := BookMarkAPI{}
	data, err := api.FindShow(context.TODO(), 32835219, "アークナイツ", 0, 48)

	assert.NoError(t, err)
	assert.NotEmpty(t, data.Works)
}

func TestBookMarkAPI_FindHide(t *testing.T) {
	api := BookMarkAPI{}
	data, err := api.FindHide(context.TODO(), 32835219, "", 0, 48)

	assert.NoError(t, err)
	assert.NotEmpty(t, data.Works)
}
