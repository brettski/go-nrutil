package main

import (
	"flag"
	"fmt"

	"github.com/brettski/go-nrutil/synthetics"
)

// GetScriptCommand is factory for getscript command
type GetScriptCommand struct {
	ID string
}

// Help is help info
func (c *GetScriptCommand) Help() string {
	helpText := `
	getscript [arg0] [arg1]
	  -id    ID of Synthetic. Overrides config file
	`

	return helpText
}

// Run is the action of getscript calling the function to get things going (running)
func (c *GetScriptCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("getscript", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		fmt.Println(c.Help())
	}
	cmdFlags.StringVar(&c.ID, "id", "", "The id of the Synthetic to get the script from. Ignores config file")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	synthetics.GetScript(c.ID)
	//fmt.Println("Running function to Get that data!", c.ID)
	return 0
}

// Synopsis is our command synopsis
func (c *GetScriptCommand) Synopsis() string {
	return "Gets the script from a synthetics scripted browser or API and stores it in a local file"
}
