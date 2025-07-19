package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeSubfolders(t *testing.T) {
	assert.ElementsMatch(t,
		[]string{"/a", "/c/d", "/c/f"},
		removeSubfolders([]string{
			"/a", "/a/b", "/c/d", "/c/d/e", "/c/f",
		}))
	assert.ElementsMatch(t,
		[]string{"/a"},
		removeSubfolders([]string{
			"/a", "/a/b/c", "/a/b/d",
		}))
	assert.ElementsMatch(t,
		[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		removeSubfolders([]string{
			"/a/b/c", "/a/b/ca", "/a/b/d",
		}))
	assert.ElementsMatch(t,
		[]string{"/a"},
		removeSubfolders([]string{
			"/a/b/c/d", "/a/b/c", "/a/b", "/a",
		}))
}
