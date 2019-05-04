package synthetics

import (
	"log"

	nrutil "github.com/brettski/go-nrutil"
)

// GetScript start process for getting a script from New Relic Synthetics
func GetScript(id string) {
	log.Println("Getting script ", id)
	config, err := nrutil.GetConfigurationInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)
}
