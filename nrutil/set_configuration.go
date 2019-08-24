package nrutil

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func createBaseYamlFile(yamlfile string) error {
	ymlDecoded, err := yaml.Marshal(&GetBaseConfiguration().DefaultConfig)
	if err != nil {
		return fmt.Errorf("Unable to marshal default config. %s", err.Error())
	}

	f, err := os.Create(yamlfile)
	if err != nil {
		return fmt.Errorf("Unable to create yaml file in home directory: %s", err.Error())
	}
	defer f.Close()
	if _, err := f.Write(ymlDecoded); err != nil {
		return fmt.Errorf("Error writing data to yaml file in home directory: %s", err.Error())
	}

	return fmt.Errorf(`
    The configuration yaml file was not in place.
    A new file was created at %s.
    Please add your New Relic User Admin key, base path and Synthetics monitor GUID's you wish to manage.
    Once set, run this again.
    `, yamlfile)
}
