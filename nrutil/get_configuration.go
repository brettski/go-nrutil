package nrutil

import (
	"encoding/base64"
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
	SyntheticMonitors []string `yaml:"syntheticmonitors"`
	BasePath          string   `yaml:"basepath"`
}

// GetConfigurationInfo reads configuration from yaml in home folder
func GetConfigurationInfo() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	yamlFile := filepath.Join(home, GetBaseConfiguration().DefaultConfigFileName)
	if _, err := os.Stat(yamlFile); err != nil {
		// file not found or other error
		if os.IsNotExist(err) {
			return nil, createBaseYamlFile(yamlFile)
		}
		return nil, err
	}

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
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
		sb.WriteString("- NrAdminKey still set to default value. Update with your NR key")
	}

	if sb.Len() > 0 {
		return fmt.Errorf("Configuration file check errors:\n%s", sb.String())
	}
	return nil
}

func createBaseYamlFile(yamlfile string) error {
	// yaml base file at /config_base.yaml
	ymldataEncoded := "LS0tCm5yYWRtaW5rZXk6IDx5b3VyLWFkbWluLWtleT4KYmFzZXBhdGg6IH4vbnJzeW50aGV0aWNzCnN5bnRoZXRpY21vbml0b3JzOgogIC0gZ3VpZC1vZi1tb25pdG9yLTEtMjM0NTYKICAtIGd1aWQtb2YtbW9uaXRvci0yLTM0NTY3CiAgLSBndWlkLW9mLW1vbml0b3Itbi1vcHFycyAK"

	ymlDecoded, err := base64.StdEncoding.DecodeString(ymldataEncoded)
	if err != nil {
		return fmt.Errorf("Unable to decode ymldata to create default file. %s", err.Error())
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
    A new one is now at %s.
    Please add you New Relic User Admin key and Synthetic monitor GUID's to manage.
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
			return nil, "", fmt.Errorf("Configuration file not found in home directory: %s", yamlFile)
		}
		return nil, "", err
	}

	data, err = ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, "", err
	}

	return
}
