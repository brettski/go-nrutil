// Package nrrequest generates our default http client for making requests
package nrrequest

import (
	"log"
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

func (r *Request) get(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating NewRequest: %s", err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	return resp, nil

	/*
		for _, option := range options {
			switch opt := option.(type) {

			case Header:
				for k, v := range opt {
					req.Header.Add(k, v)
				}

			case QueryParam:
				for k, v := range opt {
					//Add query parameter to request.
				}

			}
	*/

}
