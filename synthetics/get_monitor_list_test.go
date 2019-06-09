package synthetics

import (
	"fmt"
	"testing"
)

var monitors = []SyntheticMonitor{
	{
		Id:           "00000000-0000-0000-0000-000000000000",
		Name:         "Synthetic 0",
		MonitorType:  "SIMPLE",
		Frequency:    1,
		Uri:          "",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-05T14:19:33.292+0000",
		Modified:     "2019-06-05T14:19:33.292+0000",
		ApiVersion:   "0.5.2",
	}, {
		Id:           "11111111-1111-1111-1111-111111111111",
		Name:         "Synthetic 1",
		MonitorType:  "SCRIPT_BROWSER",
		Frequency:    5,
		Uri:          "",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-05T14:19:58.041+0000",
		Modified:     "2019-06-05T14:19:58.041+0000",
		ApiVersion:   "0.5.2",
	}, {
		Id:           "22222222-2222-2222-2222-222222222222",
		Name:         "Synthetic 2",
		MonitorType:  "SCRIPT_API",
		Frequency:    5,
		Uri:          "",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-08T15:14:03.187+0000",
		Modified:     "2019-06-08T15:14:03.187+0000",
		ApiVersion:   "0.5.2",
	}, {
		Id:           "33333333-3333-3333-3333-333333333333",
		Name:         "Synthetic 3",
		MonitorType:  "SIMPLE",
		Frequency:    1,
		Uri:          "",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-05T14:19:33.292+0000",
		Modified:     "2019-06-05T14:19:33.292+0000",
		ApiVersion:   "0.5.2",
	}, {
		Id:           "44444444-4444-4444-4444-444444444444",
		Name:         "Synthetic 4",
		MonitorType:  "SCRIPT_API",
		Frequency:    5,
		Uri:          "",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-05T14:19:58.041+0000",
		Modified:     "2019-06-05T14:19:58.041+0000",
		ApiVersion:   "0.5.2",
	}, {
		Id:           "55555555-5555-5555-5555-555555555555",
		Name:         "Synthetic 5",
		MonitorType:  "SCRIPT_API",
		Frequency:    5,
		Uri:          "",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-08T15:14:03.187+0000",
		Modified:     "2019-06-08T15:14:03.187+0000",
		ApiVersion:   "0.5.2",
	}, {
		Id:           "",
		Name:         "Synthetic -1",
		MonitorType:  "BROWSER",
		Frequency:    10,
		Uri:          "https://",
		Status:       "ENABLED",
		SlaThreshold: 1,
		Created:      "2019-06-08T15:14:03.187+0000",
		Modified:     "2019-06-08T15:14:03.187+0000",
		ApiVersion:   "0.5.2",
	},
}

var synthetics = SyntheticMonitors{
	Count:    7,
	Monitors: monitors,
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name MonitorTypeFilter
		want int
	}{
		{AnyScript, 4},
		{Browser, 1},
		{API, 3},
	}
	for _, tst := range tests {
		t.Run(string(tst.name), func(t *testing.T) {
			syntheticscopy := synthetics
			syntheticscopy.Filter(tst.name)
			if got := len(syntheticscopy.Monitors); got != tst.want {
				t.Errorf("Filter %s failed. Got: %d, Want: %d", string(tst.name), got, tst.want)
			}

		})
	}
}

func TestSortByType(t *testing.T) {
	monitors := []SyntheticMonitor{
		{MonitorType: "SCRIPTED_BROWSER", Id: "1"},
		{MonitorType: "SCRIPTED_API", Id: "2"},
		{MonitorType: "SIMPLE", Id: "0"},
		{MonitorType: "SCRIPTED_BROWSER", Id: "6"},
		{MonitorType: "SIMPLE", Id: "01"},
		{MonitorType: "SIMPLE", Id: "02"},
		{MonitorType: "BROWSER", Id: "9"},
	}
	synthetics := SyntheticMonitors{
		Count:    3,
		Monitors: monitors,
	}
	tests := []struct {
		isDecending bool
		wantfirst   string
		wantlast    string
	}{
		{true, "BROWSER", "SIMPLE"},
		{false, "SIMPLE", "BROWSER"},
	}

	for _, tst := range tests {
		t.Run(fmt.Sprintf("Sort IsDesending: %v", tst.isDecending), func(t *testing.T) {
			syntheticscopy := synthetics
			syntheticscopy.SortByType(tst.isDecending)
			gotfirst := syntheticscopy.Monitors[0].MonitorType
			gotlast := syntheticscopy.Monitors[len(syntheticscopy.Monitors)-1].MonitorType
			if tst.wantfirst != gotfirst || tst.wantlast != gotlast {
				t.Errorf("Want First: %s, Got First: %s; Want Last: %s, Got Last: %s", tst.wantfirst, gotfirst, tst.wantlast, gotlast)
			}
		})
	}
}
