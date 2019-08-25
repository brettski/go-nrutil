package synthetics

import (
	"fmt"

	termtables "github.com/brettski/go-termtables"
)

// ListMonitors shows all monitors configured in account sorted by type
func ListMonitors(showFullId bool) (string, error) {
	monitorlist, err := GetAllMonitors()
	if err != nil {
		return "", err
	}

	monitorlist.SortByType(true)

	// build table
	table := termtables.CreateTable()
	table.AddHeaders("Id", "Name", "Monitor Type", "Status")
	for _, monitor := range monitorlist.Monitors {
		var displayid string
		if showFullId {
			displayid = monitor.Id
		} else {
			displayid = monitor.Id[0:9]
		}
		table.AddRow(
			displayid,
			monitor.Name,
			monitor.MonitorType,
			monitor.Status,
		)
	}
	fmt.Println(table.Render())

	return "", nil
}
