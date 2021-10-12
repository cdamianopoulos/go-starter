package utl_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-starter/separateRepos/utl"
)

func TestHostPort(t *testing.T) {
	const longText = `Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.`
	tests := []struct {
		Host     string
		Port     uint16
		Expected string
	}{
		{"", 0, ":0"},
		{"xyz", 127, "xyz:127"},
		{longText, math.MaxUint16, longText + ":65535"},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, utl.HostPort(test.Host, test.Port))
	}
}

func TestSprint(t *testing.T) {
	tests := []struct {
		Input    []string
		Expected string
	}{
		{Input: []string{"1", "2"}, Expected: `["1", "2"]`},
		{Input: []string{"A", "B", "C", "D E F"}, Expected: `["A", "B", "C", "D E F"]`},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, utl.Sprint(test.Input))
	}
}
