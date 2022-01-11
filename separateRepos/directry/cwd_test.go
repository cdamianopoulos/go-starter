package directry_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-starter/separateRepos/directry"
)

func TestCwd(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		assert.Equal(t, ".", directry.Cwd())
	} else {
		assert.Equal(t, path, directry.Cwd())
	}
}
