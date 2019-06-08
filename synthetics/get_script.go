// Package synthetics used for interacting with synthetics api
// Performs work to get script from API.
// TODO: Break this up, all here as POC
// bubble up errors to caller, not kill here
package synthetics

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	//nrutil "github.com/brettski/go-nrutil"
	"github.com/brettski/go-nrutil/filemanager"
	"github.com/brettski/go-nrutil/nrrequest"
	"github.com/brettski/go-nrutil/nrutil"
)

// GetScript start process for getting a script from New Relic Synthetics
func GetScript(id string) {
	log.Println("Getting script ", id)
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		log.Fatal(err)
	}

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + fmt.Sprintf("monitors/%s/script", id)
	log.Println(url)

	request, _ := nrrequest.NewRequest()

	resp, err := request.Get(url)
	if err != nil {
		log.Fatalf("Request error %s", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Non-200 status code: %s\n\n", resp.Status)
		return
	}
	encodedScript, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body (json) %s\n\n", err)
		return
	}

	//log.Printf("Encoded Script:\n%s\n", encodedScript)

	var scriptPayload ScriptPayload
	if err := json.Unmarshal(encodedScript, &scriptPayload); err != nil {
		log.Fatalf("Unable to unmarshal json: %s\n\n", err)
	}

	decodedScript, err := base64.StdEncoding.DecodeString(scriptPayload.ScriptText)
	if err != nil {
		log.Fatalf("Error decoding base64 string from api: %s\n\n", err)
	}

	//log.Printf("Decoded Script:\n%s\n", decodedScript)

	fm, err := filemanager.NewFilemanager(config.BasePath, true)
	if err != nil {
		log.Fatalf("Issue creating new filemanager: %s\n\n", err)
	}

	log.Println("Writing to file")
	err = fm.WriteFile(id, decodedScript)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("file successfully written\n\n")
}
