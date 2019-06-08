package synthetics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/apcera/termtables"
	"github.com/brettski/go-nrutil/nrrequest"
	"github.com/brettski/go-nrutil/nrutil"
)

type ScriptList struct {
	Count    int32 `json:"count"`
	Monitors []struct {
		Id           string  `json:"id"`
		Name         string  `json:"name"`
		MonitorType  string  `json:"type"`
		Frequency    uint8   `json:"frequency"`
		Uri          string  `json:"uri"`
		Status       string  `json:"status"`
		SlaThreshold float32 `json:"slaThreshold"`
		Created      string  `json:"createdAt"`
		Modified     string  `json:"modifiedAt"`
		ApiVersion   string  `json"apiVersion"`
	} `json:"monitors"`
}

// ListScripts retrieves all synthetics monitors
func ListScripts() (string, error) {
	log.Println("Getting list of synthetic monitors")

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + "monitors"
	//log.Println(url)

	request, _ := nrrequest.NewRequest()

	resp, err := request.Get(url)
	if err != nil {
		return "", fmt.Errorf("Request error %s", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Non-200 status code getting data: %s", resp.Status)
	}

	respJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading body (json): %s", err)
	}

	//log.Printf("respJson returned:\n%s\n", respJson)

	var scriptList ScriptList
	if err := json.Unmarshal(respJson, &scriptList); err != nil {
		return "", fmt.Errorf("Error unmarshaling json into struct instance: %s", err)
	}
	//log.Printf("Our object: \n%+v\n", scriptList)

	// build table
	table := termtables.CreateTable()
	table.AddHeaders("Id", "Name", "Monitor Type", "Status")
	for _, monitor := range scriptList.Monitors {
		table.AddRow(
			monitor.Id[0:9],
			monitor.Name,
			monitor.MonitorType,
			monitor.Status,
		)
	}
	fmt.Println(table.Render())

	return "", nil
}
