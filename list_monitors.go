package main

import (
	"log"

	"github.com/brettski/go-nrutil/synthetics"
	cli "github.com/jawher/mow.cli"
)

func cmdListMonitors(cmd *cli.Cmd) {

	cmd.Action = func() {
		_, err := synthetics.ListMonitors()
		if err != nil {
			log.Fatalln(err)
			cli.Exit(2)
		}
	}

}
