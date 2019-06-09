package main

import (
	"log"

	"github.com/brettski/go-nrutil/synthetics"
	cli "github.com/jawher/mow.cli"
)

func cmdListMonitors(cmd *cli.Cmd) {
	fullid := cmd.BoolOpt("fullid", false, "Display full Synthetics monitor id.")

	cmd.Action = func() {
		_, err := synthetics.ListMonitors(*fullid)
		if err != nil {
			log.Println(err)
			cli.Exit(2)
		}
	}

}
