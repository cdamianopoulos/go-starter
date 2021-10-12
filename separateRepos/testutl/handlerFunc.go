package testutl

// TODO move testutl into a separate Git repository
import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// HandlerFunc sets up an HTTP ResponseRecorder request for testing.
// ok equals true when it succeeded and false when any errors occurred.
func HandlerFunc(t *testing.T, handler http.HandlerFunc, method, url string, body io.Reader, expectedStatus int) (rr *httptest.ResponseRecorder, ok bool) {
	req, err := http.NewRequest(method, url, body)
	if !assert.Nil(t, err) {
		return nil, false
	}

	rr = httptest.NewRecorder()

	// Make the handler function satisfy http.Handler
	handler.ServeHTTP(rr, req)

	return rr, assert.Equal(t, expectedStatus, rr.Code)
}
