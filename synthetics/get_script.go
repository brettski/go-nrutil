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

	"github.com/brettski/go-nrutil/filemanager"
	"github.com/brettski/go-nrutil/nrrequest"
	"github.com/brettski/go-nrutil/nrutil"
)

// GetScript start process for getting a script from New Relic Synthetics
func GetScript(id string) error {
	log.Println("Getting script ", id)
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		return err
	}

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + fmt.Sprintf("monitors/%s/script", id)
	//log.Println(url)

	request, _ := nrrequest.NewRequest()

	resp, err := request.Get(url)
	if err != nil {
		return fmt.Errorf("Request error %s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Non-200 status code: %s", resp.Status)
	}
	encodedScript, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading body (json) %s", err)
	}

	//log.Printf("Encoded Script:\n%s\n", encodedScript)

	var scriptPayload ScriptPayload
	if err := json.Unmarshal(encodedScript, &scriptPayload); err != nil {
		return fmt.Errorf("Unable to unmarshal json: %s", err)
	}

	decodedScript, err := base64.StdEncoding.DecodeString(scriptPayload.ScriptText)
	if err != nil {
		return fmt.Errorf("Error decoding base64 string from api: %s", err)
	}

	//log.Printf("Decoded Script:\n%s\n", decodedScript)

	fm, err := filemanager.NewFilemanager(config.BasePath, true)
	if err != nil {
		return fmt.Errorf("Issue creating new filemanager: %s", err)
	}

	err = fm.WriteFile(id, decodedScript)
	if err != nil {
		return fmt.Errorf("Error write to file: %s", err)
	}
	log.Printf("File successfully written\n\n")
	return nil
}
