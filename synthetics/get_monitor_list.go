package synthetics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/brettski/go-nrutil/nrrequest"
	"github.com/brettski/go-nrutil/nrutil"
)

// SyntheticMonitors represents monitor response json
type SyntheticMonitors struct {
	Count    int32              `json:"count"`
	Monitors []SyntheticMonitor `json:"monitors"`
}

// SyntheticMonitor is one monitor represented in SyneticMonitors.Monitors
type SyntheticMonitor struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	MonitorType  string  `json:"type"`
	Frequency    uint8   `json:"frequency"`
	Uri          string  `json:"uri"`
	Status       string  `json:"status"`
	SlaThreshold float32 `json:"slaThreshold"`
	Created      string  `json:"createdAt"`
	Modified     string  `json:"modifiedAt"`
	ApiVersion   string  `json:"apiVersion"`
}

// SortByType sorts SyntheticMonitors.Monitors slice
func (s SyntheticMonitors) SortByType(isDecending bool) {
	sort.Slice(s.Monitors, func(a, b int) bool {
		if isDecending {
			return s.Monitors[a].MonitorType < s.Monitors[b].MonitorType
		}
		return s.Monitors[a].MonitorType > s.Monitors[b].MonitorType
	})
}

// MonitorTypeFilter is available ways to filter SynetheicMonitors.Monitors slice
type MonitorTypeFilter string

const (
	// AnyScript defines any type of Synthetics script
	AnyScript MonitorTypeFilter = "SCRIPT"

	// Browser defines any 'Browser' type Synthetics script
	Browser MonitorTypeFilter = "SCRIPT_BROWSER"

	// API defines any 'API' type Synthetics script
	API MonitorTypeFilter = "SCRIPT_API"
)

// Filter filters current SyntheticMonitors.Monitors slice by filter type
func (s *SyntheticMonitors) Filter(f MonitorTypeFilter) {
	var filtered []SyntheticMonitor
	for _, v := range s.Monitors {
		if strings.HasPrefix(v.MonitorType, string(f)) {
			filtered = append(filtered, v)
		}
	}
	//log.Println("Lenth of filtered: ", filtered)
	if len(filtered) > 0 {
		s.Monitors = filtered
	}
}

// GetAllMonitors retrieves all Syntetic monitors from NR API
func GetAllMonitors() (*SyntheticMonitors, error) {
	log.Println("Getting list of synthetic monitors")

	baseURL := nrutil.GetBaseConfiguration().NrBaseSyntheticsAPIURL
	url := baseURL + "monitors"
	//log.Println(url)

	request, _ := nrrequest.NewRequest()

	resp, err := request.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Request error %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code getting data: %s", resp.Status)
	}

	respJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading body (json): %s", err)
	}

	//log.Printf("respJson returned:\n%s\n", respJson)

	var monitors SyntheticMonitors
	if err := json.Unmarshal(respJson, &monitors); err != nil {
		return nil, fmt.Errorf("Error unmarshaling json into struct instance: %s", err)
	}
	//log.Printf("Our object: \n%+v\n", scriptList)
	return &monitors, err

}

// SaveToConfig writes the monitors in slice, SyntheticMontiors.Monitors, to config instance. Overwrites values
func (smonitors *SyntheticMonitors) SaveToConfig(config *nrutil.Config) error {
	ids := make([]string, len(smonitors.Monitors))
	for i, mon := range smonitors.Monitors {
		ids[i] = mon.Id
	}
	config.SyntheticMonitors = ids
	return nil
}