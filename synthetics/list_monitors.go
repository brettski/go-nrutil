package synthetics

import (
	"fmt"

	"github.com/apcera/termtables"
)

// ListMonitors shows all monitors configured in account sorted by type
func ListMonitors() (string, error) {
	monitorlist, err := GetAllMonitors()
	if err != nil {
		return "", err
	}

	monitorlist.SortByType(true)

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
