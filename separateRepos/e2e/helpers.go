package e2e

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/steinfletcher/apitest"
)

type ApiError struct {
	Errs interface{} `json:"errors,omitempty"`
	Err  interface{} `json:"error,omitempty"`
}

// Get sets up a new api test to make a HTTP request
func Get(urlPattern string, a ...interface{}) *apitest.Request {
	return apitest.New().EnableNetworking(http.DefaultClient).Get(fmt.Sprintf(urlPattern, a...))
}

// Post sets up a new api test to make a HTTP request
func Post(urlPattern string, a ...interface{}) *apitest.Request {
	return apitest.New().EnableNetworking(http.DefaultClient).Post(fmt.Sprintf(urlPattern, a...))
}

// NoErrorObj returns an error when the JSON field "errors" is populated.
func NoErrorObj(rs *http.Response, _ *http.Request) error {
	data, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return fmt.Errorf("unable to read response.Body. Err: `%w`, response.Body: `%s`", err, data)
	}

	fmt.Println(string(data))

	var apiError ApiError
	// We don't care if the unmarshalling failed, only if ApiError.Err is not nil.
	_ = json.Unmarshal(data, &apiError)

	if apiError.Err != nil || apiError.Errs != nil || bytes.Contains(data, []byte("error")) {
		return fmt.Errorf("expected no errors Err: `%s`", data)
	}

	return nil
}

// HasError expects the http.Response rs to contain a JSON error object with either "error" or "errors" to be populated.
// If both of these fields are empty then HasError returns an error.
func HasError(rs *http.Response, _ *http.Request) error {
	data, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return fmt.Errorf("unable to read response.Body. Err: `%w`,\nresponse.Body: `%s`", err, data)
	}

	fmt.Println(string(data))

	if len(data) == 0 {
		return errors.New("expected the response to contain an error")
	}

	var apiError ApiError
	err = json.Unmarshal(data, &apiError)
	if err != nil {
		return fmt.Errorf("unable to unmarshal JSON into apiError as expected. Err: `%w`,\nresponse.Body: `%s`", err, data)
	}

	if apiError.Err == nil || apiError.Err == "" || apiError.Errs == nil || apiError.Errs == "" {
		return fmt.Errorf("expected the api to return at least one error. Data: `%s`", data)
	}

	return nil
}

func PrintResponse(rs *http.Response, _ *http.Request) error {
	src, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(src))
	return nil
}
