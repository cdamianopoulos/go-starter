package utl_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-starter/separateRepos/utl"
)

func TestHostPort(t *testing.T) {
	const longText = `Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old.`
	tests := []struct {
		Host     string
		Port     uint16
		Expected string
	}{
		{"", 0, ":"},
		{"", math.MaxUint8, ":255"},
		{"localhost", 0, "localhost:"},
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
		{Input: []string{"1", "2"}, Expected: "[`1`, `2`]"},
		{Input: []string{"A", "B", "C", "D E F"}, Expected: "[`A`, `B`, `C`, `D E F`]"},
		{Input: nil, Expected: "<nil>"},
		{Input: []string{}, Expected: "[]"},
		{Input: []string{""}, Expected: "[``]"},
		{Input: []string{"", ""}, Expected: "[``, ``]"},
		{Input: []string{"", "", ""}, Expected: "[``, ``, ``]"},
		{Input: make([]string, 10), Expected: "[``, ``, ``, ``, ``, ``, ``, ``, ``, ``]"},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, utl.Sprint(test.Input))
	}
}
