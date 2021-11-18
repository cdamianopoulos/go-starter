package testutl

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// JsonMarshal marshals obj, returning the JSON output as a byte slice and asserts if an error occurred.
func JsonMarshal(t *testing.T, obj interface{}) []byte {
	src, err := json.Marshal(obj)
	assert.NoError(t, err)
	return src
}

// JsonMarshalStr marshals obj, returning the JSON output as a string and asserts if an error occurred.
func JsonMarshalStr(t *testing.T, obj interface{}) string {
	return string(JsonMarshal(t, obj))
}

// JsonMarshalIndentStr marshals obj, returning the JSON output as a string and asserts if an error occurred.
func JsonMarshalIndentStr(t *testing.T, obj interface{}) string {
	src, err := json.MarshalIndent(obj, "", "\t")
	assert.NoError(t, err)
	return string(src)
}

// JsonUnmarshal unmarshalls src into obj. JsonUnmarshal expects obj to be a pointer.
func JsonUnmarshal(t *testing.T, src []byte, obj interface{}) {
	assert.NoError(t, json.Unmarshal(src, obj))
}
