package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type FileNotFoundError struct {
	filepath string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.filepath)
}

func LoadFromFile(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, FileNotFoundError{filename}
		}
		return nil, err
	}

	return Load(b)
}

func Load(bytes []byte) (*Config, error) {
	var cfgFile ConfigFile
	err := yaml.Unmarshal(bytes, &cfgFile)
	if err != nil {
		return nil, err
	}
	config, err := newFromFile(&cfgFile)
	if err != nil {
		return nil, fmt.Errorf("could not parse config file: %s", err)
	}
	return config, nil
}

func WriteToFile(config *Config, filename string) error {
	b, err := Write(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, b, 0)
	if err != nil {
		return err
	}
	return nil
}

func Write(config *Config) ([]byte, error) {
	cfgFile := config.configFile()
	b, err := yaml.Marshal(cfgFile)
	if err != nil {
		return nil, err
	}
	return b, nil
}
