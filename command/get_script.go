package main

import (
	"fmt"

	cli "github.com/jawher/mow.cli"

	"github.com/brettski/go-nrutil/synthetics"
)

func cmdGetScript(cmd *cli.Cmd) {
	id := cmd.StringOpt("id", "", "The id of the Synthetic to get the script from. Ignores config file")

	cmd.Action = func() {
		if len(*id) < 1 {
			fmt.Printf("id is a required parameter")
			cli.Exit(1)
		}
		fmt.Println("Do get script stuff", id)
		synthetics.GetScript(*id)
	}
}
