// Package synthetics used for interacting with synthetics api
// Performs work to get script from API.
// TODO: Break this up, all here as POC
// bubble up errors to caller, not kill here
package synthetics

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/brettski/go-nrutil/filemanager"
	"github.com/brettski/go-nrutil/nrrequest"
	"github.com/brettski/go-nrutil/nrutil"
)

// SetScript processes writing a script from file to New Relic Synthetics
func SetScript(id string) error {
	log.Println("Setting script ", id)
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		return err
	}

	fm, err := filemanager.NewFilemanager(config.BasePath, false)
	if err != nil {
		return fmt.Errorf("Error creating new filemanager: %s", err)
	}

	decodedScript, err := fm.ReadFile(id)
	if err != nil {
		return fmt.Errorf("Error reading local script file: %s", err)
	}

	encodedScript := base64.StdEncoding.EncodeToString(decodedScript)
	scriptPayload := &ScriptPayload{
		ScriptText: encodedScript,
	}

	payload, err := json.Marshal(&scriptPayload)
	if err != nil {
		return fmt.Errorf("Error marshaling payload into struct: %s", err)
	}

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + fmt.Sprintf("monitors/%s/script", id)
	//log.Println(url)

	request, _ := nrrequest.NewRequest()
	resp, err := request.Put(url, payload)
	if err != nil {
		return fmt.Errorf("Request error %s", err)
	}

	if resp.StatusCode != 204 {
		//log.Printf("whole resp: %+v", resp)
		return fmt.Errorf("Put request was not succesful. Status: %s", resp.Status)
	}

	log.Printf("file %s successfuly sent to New Relic\n\n", id)
	return nil
}
