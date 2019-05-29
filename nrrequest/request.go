// Package nrrequest generates our default http client for making requests
package nrrequest

import (
	"net/http"
)

// Request struct is a simple wrapper around the net/http library
// map exmpl-https://play.golang.com/p/OW8FSpiD-9k
type Request struct {
	client     *http.Client
	req        *http.Request
	Header     Header
	QueryParam QueryParam
	URL        string
}

// Header type for setting header options for the request.
type Header map[string]string

// QueryParam type for setting query parameters for the request
type QueryParam map[string]string

// NewRequest generates a new Request object
func NewRequest() (*Request, error) {

	return &Request{
		client: &http.Client{},
	}, nil

}
