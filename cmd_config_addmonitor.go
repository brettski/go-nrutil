package main

import (
	"log"

	"github.com/brettski/go-nrutil/synthetics"
	"github.com/brettski/go-nrutil/nrutil"
	cli "github.com/jawher/mow.cli"
)

func cmdAddMonitor(cmd *cli.Cmd) {
	id := cmd.StringOpt("id", "", "Add 1 id to configuration file")
	all := cmd.BoolOpt("a all", false, "Add all SCRIPT type Synthetic monitors")

	cmd.Action = func() {
		if len(*id) > 0 && *all {
			log.Println("Only use id OR all option. Using both is not valid")
			cli.Exit(1)
		}
		if len(*id) < 1 && !*all {
			log.Println("No options provided, there is nothing to do. Exiting")
			cli.Exit(1)
		}
		if *all {
			configuration, err := nrutil.GetConfigurationInfo()
			if err != nil {
				log.Printf("Unable to retrieve config file.\n%s\n", err)
				cli.Exit(2)
			}
			monitors, err := synthetics.GetAllMonitors()
			if err != nil {
				log.Printf("Unable to retrieve Synthetics Monitor list from New Relic.\n%s\n", err)
				cli.Exit(2)
			}
			monitors.Filter(synthetics.AnyScript)

			log.Printf("Before synthetics count: %d", len(configuration.SyntheticMonitors))
			err = monitors.SaveToConfig(configuration)
			if err!= nil {
				log.Printf("Error occured updating configuration with all Synthetics id's.\n%s\n", err)
				cli.Exit(2)
			}
			log.Printf("After  synthetics count: %d", len(configuration.SyntheticMonitors))
			err = configuration.SetConfigurationInfo(nrutil.GetConfigurationFileLocation())
			if err != nil {
				log.Printf("Unable to save configuration file.\n%s\n", err)
				cli.Exit(2)
			}

		} else {
			log.Println("No action taken")
		}

	}
}
