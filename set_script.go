package main

import (
	"fmt"

	"github.com/brettski/go-nrutil/synthetics"
	cli "github.com/jawher/mow.cli"
)

func cmdSetScript(cmd *cli.Cmd) {
	id := cmd.StringOpt("id", "", "The id of the Synthetic to get the script from. Ignores config file")

	cmd.Action = func() {
		if len(*id) < 1 {
			fmt.Printf("id is a requireed parameter\n\n")
			cli.Exit(1)
		}
		fmt.Println("Do set	script stuff")
		synthetics.SetScript(*id)
	}
}
