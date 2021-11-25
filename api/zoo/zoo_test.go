package zoo_test

import (
	"go-starter/api/zoo"
	"go-starter/separateRepos/testutl"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert.Equal(t, "/zoo", zoo.URLPattern)
}

func TestHandler(t *testing.T) {
	rr, ok := testutl.HandlerFunc(t, zoo.Status, http.MethodGet, zoo.URLPattern, "", http.StatusOK)
	assert.True(t, ok)

	var response struct {
		Zoo       string
		Visitors  int
		AnimalQty int `json:"animal_qty"`
	}
	testutl.JsonUnmarshal(t, rr.Body.Bytes(), &response)

	assert.Equal(t, "closed", response.Zoo)
	assert.Equal(t, 11, response.AnimalQty)
	assert.Condition(t,
		func() bool {
			return response.Visitors >= 0 && response.Visitors <= 59
		},
		"visitors must be an integer between 0 and 59",
		response.Visitors,
	)
}
