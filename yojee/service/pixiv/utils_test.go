package pixiv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_Path(t *testing.T) {
	assert.Equal(t, Path("/ajax", "follow_latest", "illust"), "/ajax/follow_latest/illust")
	assert.Equal(t, Path("ajax", "follow_latest", "illust"), "/ajax/follow_latest/illust")
	assert.Equal(t, Path("ajax", "follow_latest", 1), "/ajax/follow_latest/1")
	assert.Equal(t, Path("ajax", "follow_latest", 1, 2), "/ajax/follow_latest/1/2")
}
