package nrutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	yaml "gopkg.in/yaml.v2"
)

// Config is configuration struct read from yaml
type Config struct {
	NrAdminKey        string   `yaml:"nradminkey"`
	SyntheticMonitors []string `yaml:"syntheticmonitors"`
}

// GetConfigurationInfo reads configuration from yaml in home folder
func GetConfigurationInfo() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	yamlFile := filepath.Join(home, ".nrutil.yml")
	if _, err := os.Stat(yamlFile); err != nil {
		// file not found
		return nil, createBaseYamlFile(yamlFile)
	}

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	nrconfig := &Config{}
	if err := yaml.Unmarshal(data, nrconfig); err != nil {
		return nil, err
	}

	return nrconfig, nil
}

// Check verifies that configuration has expected, needed fields
func (c *Config) Check() error{
	// Not sure if this is the best approach, but lets see
	var errs string
	if len(c.NrAdminKey) < 10 {
		errs += fmt.Sprintf("NrAdminKey not set or too short\n")
	}

	if len(errs) > 0 {
		return fmt.Errorf("Configuration file check errors:\n%s", errs)
	}
	return nil
}

func createBaseYamlFile(yamlfile string) error {
	ymldata := []byte("---\nnradminkey: <your-admin-key>\nsyntheticmonitors:\n  - guid-of-monitor-1-23456\n  - guid-of-monitor-2-34567\n  - guid-of-monitor-n-opqrs")

	f, err := os.Create(yamlfile)
	if err != nil {
		return fmt.Errorf("Unable to create yaml file in home directory: %s", err.Error())
	}
	defer f.Close()
	if _, err := f.Write(ymldata); err != nil {
		return fmt.Errorf("Error writing data to yaml file in home directory: %s", err.Error())
	}

	return fmt.Errorf(`
    The configuration yaml file was not in place.
    A new one is now at %s.
    Please add you New Relic User Admin key and Synthetic monitor GUID's to manage.
    Once set, run this again.
    `, yamlfile)
}
