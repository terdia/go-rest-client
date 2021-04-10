package mocks

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	DoRequestFunc func(req *http.Request) (*http.Response, error)
)

type MockHttpClient struct {
}

type MockResponse struct {
	StatusCode int    // e.g. 200
	Body       string // "All good"
}

func (mock *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return DoRequestFunc(req)
}

func NewMockResponse(resp *MockResponse) {

	if resp == nil || resp.Body == "" {
		DoRequestFunc = func(*http.Request) (*http.Response, error) {
			return nil, nil
		}
		return
	}

	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(resp.Body)))

	response := &http.Response{
		StatusCode: resp.StatusCode,
		Body:       responseBody,
	}

	DoRequestFunc = func(*http.Request) (*http.Response, error) {
		return response, nil
	}
}

func NewMockError(message string) {
	DoRequestFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(message)
	}
}
