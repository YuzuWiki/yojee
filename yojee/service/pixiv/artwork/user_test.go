package artwork

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAPI_UserProfile(t *testing.T) {
	ctx := context.Background()

	u := UserAPI{}
	userProfile, err := u.UserProfile(ctx, 39123643)

	require.NoError(t, err)
	assert.NotEmpty(t, userProfile)
}
