package cachetemplates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const templateFolder = "../templates"

func TestCacher(t *testing.T) {
	c, errNew := NewCacher(templateFolder)
	require.Nil(t, errNew, "could not create cacher")
	require.Greater(t, len(c), 0)
}

func TestCacherSlash(t *testing.T) {
	c, errNew := NewCacher(templateFolderSlash)
	require.Nil(t, errNew, "could not create cacher")
	require.Greater(t, len(c), 0)
}
