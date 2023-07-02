package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

func ParseFile(name string) (*Config, error) {
	err := createDefaultFileIfNotExists(name)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	conf := &Config{}

	err = toml.NewDecoder(file).Decode(conf)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	err = validate(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func createDefaultFileIfNotExists(name string) error {
	_, err := os.Stat(name)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		conf := getDefaultConfig()
		file, err := os.Create(name)
		if err != nil {
			return err
		}

		enc := toml.NewEncoder(file)
		enc.SetIndentTables(false)
		enc.SetArraysMultiline(true)
		err = enc.Encode(conf)
		if err != nil {
			return err
		}

		return file.Close()
	}
	return nil
}
