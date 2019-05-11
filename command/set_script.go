package main

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
)

func cmdSetScript(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("Do set	cript stuff")
	}
}
