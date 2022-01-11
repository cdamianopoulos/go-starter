package healthcheck_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-starter/separateRepos/healthcheck"
	"go-starter/separateRepos/testutl"
)

func TestURL(t *testing.T) {
	assert.Equal(t, "/status", healthcheck.URLPattern)
}

func TestHandler(t *testing.T) {
	rr, ok := testutl.HandlerFunc(t, healthcheck.Status, http.MethodGet, healthcheck.URLPattern, "", http.StatusOK)

	assert.Equal(t, `{"status":"healthy"}`, rr.Body.String())
	assert.True(t, ok)
}
