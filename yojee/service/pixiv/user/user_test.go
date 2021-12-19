package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAPI_Extra(t *testing.T) {
	api := UsersAPI{}

	data, err := api.Extra(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Following)
}
