// Package nrrequest generates our default http client for making requests
package nrrequest

import (
	"fmt"
	"net/http"

	"github.com/brettski/go-nrutil/nrutil"
)

// Get performs a GET request for provided URL
func (r *Request) Get(url string) (*http.Response, error) {

	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating NewRequest: %s", err)
	}
	req.Header.Add("X-Api-Key", config.NrAdminKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %s", err)
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
