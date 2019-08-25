package main

import (
	"fmt"
	"log"

	"github.com/brettski/go-nrutil/synthetics"
	cli "github.com/jawher/mow.cli"
)

func cmdSetScript(cmd *cli.Cmd) {
	id := cmd.StringOpt("id", "", "A specific id of a Synthetic monitor to set a script to.")

	cmd.Action = func() {
		if len(*id) < 1 {
			fmt.Printf("id is a requireed parameter\n\n")
			cli.Exit(1)
		}
		fmt.Println("Do 'set	script' stuff")
		err := synthetics.SetScript(*id)
		if err != nil {
			log.Println(err)
			cli.Exit(2)
		}
	}
}
