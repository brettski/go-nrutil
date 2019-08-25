package main

import (
	"log"

	cli "github.com/jawher/mow.cli"
)

func cmdAddMonitor(cmd *cli.Cmd) {
	id := cmd.StringOpt("id", "", "Add 1 id to configuration file")
	all := cmd.BoolOpt("a all", false, "Add all SCRIPT type Synthetic monitors")

	cmd.Action = func() {
		if len(*id) > 0 && *all {
			log.Println("Only use id OR all option. Using both is not valid")
			cli.Exit(1)
		}
		if len(*id) < 1 && !*all {
			log.Println("No options provided, there is nothing to do. Exiting")
			cli.Exit(1)
		}

	}
}
