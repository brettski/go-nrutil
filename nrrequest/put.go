// Package nrrequest generates our default http client for making requests
package nrrequest

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/brettski/go-nrutil/nrutil"
)

// Put performs a PUT request for the provided nr api url
func (r *Request) Put(url string, data []byte) (*http.Response, error) {
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Error creating newRequest: %s", err)
	}

	req.Header.Add("X-Api-Key", config.NrAdminKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %s", err)
	}

	return resp, nil
}
