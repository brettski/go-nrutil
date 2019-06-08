package main

import (
	"log"

	"github.com/brettski/go-nrutil/synthetics"
	cli "github.com/jawher/mow.cli"
)

func cmdGetScript(cmd *cli.Cmd) {
	id := cmd.StringOpt("id", "", "A specfic id of a Synthetic monitor to get a script from. Using this overrides config file")

	cmd.Action = func() {
		if len(*id) < 1 {
			log.Printf("id is a required parameter\n\n")
			cli.Exit(1)
		}
		log.Println("Do 'get script' stuff")
		err := synthetics.GetScript(*id)
		if err != nil {
			log.Fatalln(err)
			cli.Exit(2)
		}
	}
}
