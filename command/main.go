package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {

	c := cli.NewCLI("nrutil", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"getscript": func() (cli.Command, error) {
			return &GetScriptCommand{}, nil
		},
		"setscript": func() (cli.Command, error) {
			return &SetScriptCommand{}, nil
		},
	}

	exitStatus, err := c.Run()

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
