package nrutil

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func createBaseYamlFile(yamlfile string) error {

	if err := setConfigurationInfo(&GetBaseConfiguration().DefaultConfig, yamlfile); err != nil {
		return fmt.Errorf("Unable to marshal default config. %s", err.Error())
	}
	return fmt.Errorf(`
    The configuration yaml file was not in place.
    A new file was created at %s.
    Please add your New Relic User Admin key, base path and Synthetics monitor GUID's you wish to manage.
    Once set, run this again.
    `, yamlfile)
}

func setConfigurationInfo(configuration *Config, yamlfile string) error {
	yamlBytes, err := yaml.Marshal(&configuration)
	if err != nil {
		return fmt.Errorf("Unable to marshal Config object to yaml byte array: %s", err.Error())
	}

	// Open the file, create if it doesn't exist. Overwrite current contents

	f, err := os.OpenFile(yamlfile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Error opening or creating configuration file: %s", err.Error())
	}
	defer f.Close()
	if _, err = f.Write(yamlBytes); err != nil {
		return fmt.Errorf("Error writing to configuration file: %s", err.Error())
	}
	return nil
}
