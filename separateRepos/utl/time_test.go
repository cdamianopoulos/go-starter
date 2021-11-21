package utl

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMilliseconds(t *testing.T) {
	location, err := time.LoadLocation("GMT")
	assert.NoError(t, err)

	for i := 0; i <= 999; i++ {
		assert.Equal(t, i, Milliseconds(
			time.Date(0, 0, 0, 0, 0, 0, i*int(time.Millisecond), location),
		))
	}
}
