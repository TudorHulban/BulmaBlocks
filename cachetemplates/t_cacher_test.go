package cachetemplates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const templateFolder = "../templates"

func TestCacher(t *testing.T) {
	c, errNew := NewCacher(templateFolder)
	require.Nil(t, errNew, "could not create cacher")
	require.Empty(t, c)
}
