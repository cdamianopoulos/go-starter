package testutl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// HandlerFunc sets up an HTTP ResponseRecorder request for testing.
// urlPattern uses fmt.Sprintf to format according to a format specifier.
// method is one of the HTTP Methods.
// requestBody is the payload to send.
// ok returns true when the ResponseRecorder succeeded and the returned HTTP status matches statusCode,
// otherwise false if any errors occurred.
func HandlerFunc(t *testing.T, handler http.HandlerFunc, method, urlPattern, requestBody string, expectedStatus int, a ...interface{}) (rr *httptest.ResponseRecorder, ok bool) {
	if len(a) >= 1 && a[0] != nil {
		urlPattern = fmt.Sprintf(urlPattern, a...)
	}
	req, err := http.NewRequest(method, urlPattern, strings.NewReader(requestBody))
	// Assert there was no error.
	if !assert.NoError(t, err) {
		// But if there was an error, then return.
		return nil, false
	}

	rr = httptest.NewRecorder()

	// Make the handler function satisfy http.Handler.
	handler.ServeHTTP(rr, req)

	return rr, assert.Equal(t, expectedStatus, rr.Result().StatusCode)
}

// HandlerFuncBody is identical to HandlerFunc, but also returns the HTTP response body as a string.
func HandlerFuncBody(t *testing.T, handler http.HandlerFunc, method, urlPattern, requestBody string, expectedStatus int, a ...interface{}) (rr *httptest.ResponseRecorder, responseBody string, ok bool) {
	rr, ok = HandlerFunc(t, handler, method, urlPattern, requestBody, expectedStatus, a...)
	if rr != nil && rr.Body != nil {
		return rr, rr.Body.String(), ok
	}

	return
}
