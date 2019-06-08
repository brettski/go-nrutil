package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("nrutil", "New Relic Synthetics utility")
	app.Version("v version", "0.1.0")
	app.Command("getscript", "Retrieves script from Synthetic monitor to local file", cmdGetScript)
	app.Command("setscript", "Saves script from local file to Synthetic monitor", cmdSetScript)

	app.Run(os.Args)

}
