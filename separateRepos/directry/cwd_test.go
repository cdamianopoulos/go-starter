package directry_test

import (
	"go-starter/separateRepos/directry"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCwd(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		assert.Equal(t, ".", directry.Cwd())
	} else {
		assert.Equal(t, path, directry.Cwd())
	}
}
