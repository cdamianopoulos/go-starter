package dbconn

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	message = "input: `%d`"
	maxQty  = 1<<maxBitSize - 1
)

func TestMaxQty(t *testing.T) {
	assert.Equal(t, 4095, maxQty) // Remember to update comments whenever maxBitSize is changed.
}

func TestOpenQty_Set(t *testing.T) {
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		var iq OpenQty
		err := iq.Set(strconv.Itoa(i))

		if i >= 0 && i <= maxQty {
			// Test a valid number is assigned and no error is returned.
			if !assert.Nil(t, err, message, i) || !assert.Equal(t, i, int(iq), message, i) {
				return
			}
		} else {
			// Test nothing is assigned and an error is returned.
			if !assert.NotNil(t, err, message, i) || !assert.Equal(t, 0, int(iq), message, i) {
				return
			}
		}
	}
}

func TestIdleQty_Set(t *testing.T) {
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		var iq IdleQty
		err := iq.Set(strconv.Itoa(i))
		if i >= -1 && i <= maxQty {
			// Test a valid number is assigned and no error is returned.
			if !assert.Nil(t, err, message, i) || !assert.Equal(t, i, int(iq), message, i) {
				return
			}
		} else {
			// Test nothing is assigned and an error is returned.
			if !assert.NotNil(t, err, message, i) || !assert.Equal(t, 0, int(iq), message, i) {
				return
			}
		}
	}
}
