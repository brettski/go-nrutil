package main

import (
	"log"

	"github.com/brettski/go-nrutil/synthetics"
	cli "github.com/jawher/mow.cli"
)

func cmdListScripts(cmd *cli.Cmd) {
	cmd.Action = func() {
		_, err := synthetics.ListScripts()
		if err != nil {
			log.Fatalln(err)
			cli.Exit(2)
		}
	}

}
