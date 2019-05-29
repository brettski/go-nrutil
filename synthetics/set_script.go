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

	nrutil "github.com/brettski/go-nrutil"
	"github.com/brettski/go-nrutil/filemanager"
	"github.com/brettski/go-nrutil/nrrequest"
)

// SetScript processes writing a script from file to New Relic Synthetics
func SetScript(id string) {
	log.Println("setting script ", id)
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		log.Fatalln(err)
	}

	fm, err := filemanager.NewFilemanager(config.BasePath, false)
	if err != nil {
		log.Fatalf("Error creating new filemanager: %s\n", err)
	}

	decodedScript, err := fm.ReadFile(id)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	encodedScript := base64.StdEncoding.EncodeToString(decodedScript)
	scriptPayload := &ScriptPayload{
		ScriptText: encodedScript,
	}

	payload, err := json.Marshal(&scriptPayload)
	if err != nil {
		log.Fatalf("%s", err)
	}

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + fmt.Sprintf("monitors/%s/script", id)
	log.Println(url)

	request, _ := nrrequest.NewRequest()
	resp, err := request.Put(url, payload)
	if err != nil {
		log.Fatalf("request error %s\n", err)
	}

	if resp.StatusCode != 204 {
		//log.Printf("whole resp: %+v", resp)
		log.Fatalf("Put request was not succesful. Status: %s", resp.Status)
	} else {
		log.Printf("file %s successfuly sent to New Relic", id)
	}

}
