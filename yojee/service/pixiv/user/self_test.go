package user

import (
	"context"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestSelfAPI_Extra(t *testing.T) {
	api := SelfAPI{}

	data, err := api.Extra(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Following)
}
