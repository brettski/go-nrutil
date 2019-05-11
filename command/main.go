package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("nrutil", "New Relic Synthetics utility")
	app.Command("getscript", "Retrieves script from Synthetic monitor to local file", cmdGetScript)
	app.Command("setscript", "Saves script from local file to Synthetic monitor", cmdSetScript)

	app.Run(os.Args)

	//c := cli.NewCLI("nrutil", "0.1.0")
	//c.Args = os.Args[1:]
	/*
		c.Commands = map[string]cli.CommandFactory{
			"getscript": func() (cli.Command, error) {
				return &GetScriptCommand{}, nil
			},
			"setscript": func() (cli.Command, error) {
				return &SetScriptCommand{}, nil
			},
		}
	*/
	//exitStatus, err := c.Run()

}
