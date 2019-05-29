// Package nrrequest generates our default http client for making requests
package nrrequest

import (
	"bytes"
	"log"
	"net/http"

	"github.com/brettski/go-nrutil"
)

// Put performs a PUT request for the provided nr api url
func (r *Request) Put(url string, data []byte) (*http.Response, error) {
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating newRequest: %s\n", err)
	}

	req.Header.Add("X-Api-Key", config.NrAdminKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s\n", err)
		return nil, err
	}

	return resp, nil
}
