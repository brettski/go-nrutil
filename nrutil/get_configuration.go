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

// Config is configuration struct read from yaml
type Config struct {
	NrAdminKey        string   `yaml:"nradminkey"`
	BasePath          string   `yaml:"basepath"`
	SyntheticMonitors []string `yaml:"syntheticmonitors"`
}

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
