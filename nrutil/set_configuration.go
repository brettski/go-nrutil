package nrutil

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CreateBaseConfigFile creates a new base configuration file in $HOME folder
func CreateBaseConfigFile(NrAdminKey string) error {

	baseConfig := GetBaseConfiguration().DefaultConfig
	yamlfile := GetConfigurationFileLocation()

	if len(NrAdminKey) > 1 {
		baseConfig.NrAdminKey = NrAdminKey
	}

	if err := baseConfig.SetConfigurationInfo(yamlfile); err != nil {
		return fmt.Errorf("Unable to marshal default config. %s", err.Error())
	}
	fmt.Printf("New base configuration file set at: %s\n\n", yamlfile)
	return nil
}

// SafeCreateBaseConfigFile safely creates configuration file only overwriting if Force is true
func SafeCreateBaseConfigFile(Force bool, NrAdminKey string) error {
	configfile := GetConfigurationFileLocation()
	FileInfo, err := os.Stat(configfile)
	if err != nil {
		if os.IsNotExist(err) {
			return CreateBaseConfigFile(NrAdminKey)
		}
		return err
	}
	if FileInfo != nil && Force { // a file is found no error returned from stat()
		return CreateBaseConfigFile(NrAdminKey)
	}
	return fmt.Errorf("Configuration file already exists. Use --force flag to overwrite")

	/*
		fmt.Printf("Force %v\n", Force)
		configfile = GetConfigurationFileLocation()
		finfo, err = os.Stat(configfile)
		fmt.Printf("locat %s\n", configfile)
		fmt.Printf("Exist %v\n", os.IsExist(err))
		fmt.Printf("noexi %v\n", os.IsNotExist(err))
		fmt.Printf("Info: %+v\n", finfo)
	*/
}

// SetConfigurationInfo writes out configuration file from provided config struct instance
func (configuration *Config) SetConfigurationInfo(yamlfile string) error {
	yamlBytes, err := yaml.Marshal(&configuration)
	if err != nil {
		return fmt.Errorf("Unable to marshal Config object to yaml byte array: %s", err.Error())
	}

	// Open the file, create if it doesn't exist. Overwrite current contents

	f, err := os.OpenFile(yamlfile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error opening or creating configuration file: %s", err.Error())
	}
	defer f.Close()
	if _, err = f.Write(yamlBytes); err != nil {
		return fmt.Errorf("Error writing to configuration file: %s", err.Error())
	}
	return nil
}
