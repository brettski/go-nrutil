package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("nrutil", "New Relic Synthetics Utility")
	app.Version("v version", "0.1.0")
	app.Command("getscript", "Retrieves script from Synthetics monitor to local file", cmdGetScript)
	app.Command("setscript", "Saves script from local file to Synthetics monitor", cmdSetScript)
	app.Command("listmonitors", "Retrieves a list of all synthetic monitors in API key's account", cmdListMonitors)
	app.Command("config", "Configuration actions for CLI", cmdConfigActions)

	app.Run(os.Args)

}
