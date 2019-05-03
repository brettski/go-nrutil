package main

import "fmt"

// SetScriptCommand is the setscript command struct
type SetScriptCommand struct {
	ID string
}

// Help text for setscript
func (t *SetScriptCommand) Help() string {
	return "hello [arg0] [arg1] ... says hello to everyone"
}

// Run will do the work for setscript
func (t *SetScriptCommand) Run(args []string) int {
	fmt.Println("hello", args)
	return 0
}

// Synopsis is information about setscript
func (t *SetScriptCommand) Synopsis() string {
	return "A sample command that says hello on stdout"
}
