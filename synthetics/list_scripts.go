package synthetics

import (
	"fmt"
	"sort"

	"github.com/apcera/termtables"
)

// ListScripts retrieves all synthetics monitors
func ListScripts() (string, error) {
	monitorlist, err := GetAllMonitors()
	if err != nil {
		return "", err
	}

	sort.Slice(monitorlist.Monitors, func(a, b int) bool {
		return monitorlist.Monitors[a].MonitorType < monitorlist.Monitors[b].MonitorType
	})

	// build table
	table := termtables.CreateTable()
	table.AddHeaders("Id", "Name", "Monitor Type", "Status")
	for _, monitor := range monitorlist.Monitors {
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
