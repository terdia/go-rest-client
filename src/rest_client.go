package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var (
	Client ClientInterface
)

type ClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{
		Timeout: time.Second * 10,
	}
}

// make an HTTP post request, with data(STRUCT) to given URL
func Post(url string, data interface{}, headers http.Header) (*http.Response, error) {
	dataJsonBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(dataJsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return Client.Do(req)
}

func Get(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return Client.Do(req)
}

func Patch(url string, data interface{}, headers http.Header) (*http.Response, error) {
	dataJsonBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(dataJsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return Client.Do(req)
}

func Put(url string, data interface{}, headers http.Header) (*http.Response, error) {
	dataJsonBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(dataJsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return Client.Do(req)
}

func Delete(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return Client.Do(req)
}
