package nrutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	yaml "gopkg.in/yaml.v2"
)

// GetConfigurationInfo reads configuration from yaml in home folder
func GetConfigurationInfo() (*Config, error) {

	data, yamlFile, err := GetConfigurationFile()
	// file not found or other error
	if os.IsNotExist(err) {
		return nil, createBaseYamlFile(yamlFile)
	}

	nrconfig := &Config{}
	if err := yaml.Unmarshal(data, nrconfig); err != nil {
		return nil, err
	}

	if err := nrconfig.check(); err != nil {
		return nil, err
	}

	return nrconfig, nil
}

// Check verifies that configuration has expected, needed fields
func (c *Config) check() error {
	var sb strings.Builder
	if len(c.NrAdminKey) < 10 {
		sb.WriteString("- NrAdminKey not set or too short\n")
	}

	if strings.HasPrefix(c.NrAdminKey, "<") {
		sb.WriteString("- NrAdminKey still set to default value. Update with your NR Admin API key")
	}

	if sb.Len() > 0 {
		return fmt.Errorf("Configuration file check errors:\n%s", sb.String())
	}
	return nil
}

// GetConfigurationFile opens configuration file and returns a byte slice of contents, or error if not exist
func GetConfigurationFile() (data []byte, yamlFile string, err error) {
	var (
		home string
	)

	home, err = homedir.Dir()
	if err != nil {
		return nil, "", err
	}

	yamlFile = filepath.Join(home, GetBaseConfiguration().DefaultConfigFileName)
	if _, err = os.Stat(yamlFile); err != nil {
		if os.IsNotExist(err) {
			return nil, yamlFile, err
		}
		return nil, yamlFile, err
	}

	data, err = ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, yamlFile, err
	}

	return
}
