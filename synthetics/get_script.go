// Package synthetics used for interacting with synthetics api
// Performs work to get script from API.
// TODO: Break this up, all here as POC
// bubble up errors to caller, not kill here
package synthetics

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	nrutil "github.com/brettski/go-nrutil"
)

// GetScript start process for getting a script from New Relic Synthetics
func GetScript(id string) {
	log.Println("Getting script ", id)
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		log.Fatal(err)
	}
	if err := config.Check(); err != nil {
		log.Fatal(err)
	}

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + fmt.Sprintf("monitors/%s/script", id)
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error settin gup request:" + err.Error())
		return
	}
	req.Header.Add("X-Api-Key", config.NrAdminKey)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request error: ", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		log.Fatal("Non-200 status code: ", resp.StatusCode, resp.Status)
		return
	}
	encodedScript, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body (json)" + err.Error())
		return
	}

	log.Printf("Encoded Script:\n%s\n", encodedScript)

}
