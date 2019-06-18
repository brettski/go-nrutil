package main

import (
	"log"

	"github.com/brettski/go-nrutil/nrutil"
	cli "github.com/jawher/mow.cli"
)

func cmdConfigActions(cmd *cli.Cmd) {

	cmd.Command("dump", "Writes the contents of the current configuration file to stdout", cmdDump)
	cmd.Command("addmonitor", "Add Synthetic monitors to configuration file", cmdAddMonitor)
	cmd.Command("create", "Creates a new configuration file in your home folder.", cmdCreateConfigFile)
	cmd.Command("dump", "Writes the contents of the current configuration file to stdout", cmdDump)
	cmd.Command("setvalue", "Set non array values of config like api key and base path", cmdSetValue)

}

func cmdDump(cmd *cli.Cmd) {

	cmd.Action = func() {
		data, file, err := nrutil.GetConfigurationFile()
		if err != nil {
			log.Println("Unable to display configuration file. Reason: ", err)
			cli.Exit(3)
		}
		log.Printf("Configuration file: %s\n", file)
		log.Printf("\n\n%s\n\n", string(data))
	}
}

func cmdCreateConfigFile(cmd *cli.Cmd) {
	//force := cmd.BoolOpt("f force", false, "Force overwrite if file already exists.")
	cmd.Action = func() {

	}
}

/*
 * Need to work through management of the config file.
 * We need to abstract out the work for reuse
 * There are two created stages, one or two sets, etc. Break all that out.
 * figure out at what point we write back out to the file. :)
 */
