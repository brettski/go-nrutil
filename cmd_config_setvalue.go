package main

import (
	"log"

	cli "github.com/jawher/mow.cli"
)

func cmdSetValue(cmd *cli.Cmd) {
	apikey := cmd.StringOpt("apikey", "", "Set User Admin API key. Account REST API key will not work")
	basepath := cmd.StringOpt("basepath", "", "Set base file path to store New Relic Synthetics scripts")

	cmd.Action = func() {
		if len(*apikey) < 0 && len(*basepath) < 0 {
			log.Println("No option provided, nothing to update. Use -h for help")
			cli.Exit(1)
		}
	}
}
