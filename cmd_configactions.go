package main

import (
	"log"

	"github.com/brettski/go-nrutil/nrutil"
	cli "github.com/jawher/mow.cli"
)

func cmdConfigActions(cmd *cli.Cmd) {

	cmd.Command("dump", "Writes the contents of the current configuration file to stdout", cmdDump)
	cmd.Command("addmonitor", "Add Synthetic monitors to configuration file", cmdAddMonitor)
	cmd.Command("setvalue", "Set non array values of config like api key and base path", cmdSetValue)

}

func cmdDump(cmd *cli.Cmd) {

	cmd.Action = func() {
		data, file, err := nrutil.GetConfigurationFile()
		if err != nil {
			log.Println("Unable to display configuration file. Reason: ", err)
			cli.Exit(3)
		}
		log.Printf("Configuration file: %s\n", file)
		log.Printf("\n\n%s\n\n", string(data))
	}
}
