package testutl_test

import (
	"go-starter/separateRepos/testutl"
	"testing"

	"github.com/stretchr/testify/assert"
)

type animal struct {
	name, species string
}

type public struct {
	Name, Species string
}

func TestJsonMarshalIndentStr(t *testing.T) {
	const exportedStruct = `{
	"Name": "bob",
	"Species": "zebra"
}`

	tests := []struct {
		name     string
		expected string
		obj      interface{}
	}{
		{name: "nil", expected: `null`, obj: nil},
		{name: "struct", expected: `{}`, obj: animal{
			name:    "bob",
			species: "zebra",
		}},
		{name: "exported struct", expected: exportedStruct, obj: public{
			Name:    "bob",
			Species: "zebra",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testutl.JsonMarshalIndentStr(t, tt.obj); got != tt.expected {
				t.Errorf("JsonMarshalIndentStr() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestJsonMarshalStr(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		obj      interface{}
	}{
		{name: "nil", expected: `null`, obj: nil},
		{name: "struct", expected: `{}`, obj: animal{
			name:    "bob",
			species: "zebra",
		}},
		{name: "exported struct", expected: `{"Name":"bob","Species":"zebra"}`, obj: public{
			Name:    "bob",
			Species: "zebra",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, testutl.JsonMarshalStr(t, tt.obj), tt.expected)
		})
	}
}

func TestJsonUnmarshal(t *testing.T) {
	t.Run("unmarshal animal", func(t *testing.T) {
		var p animal
		testutl.JsonUnmarshal(t, []byte(`{"name":"bob","species":"zebra"}`), &p)
		assert.Equal(t, p, animal{name: "", species: ""})
	})

	t.Run("unmarshal public", func(t *testing.T) {
		var p public
		testutl.JsonUnmarshal(t, []byte(`{"name":"bob","species":"zebra"}`), &p)
		assert.Equal(t, p, public{Name: "bob", Species: "zebra"})
	})
}
