package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	PicturePath string `toml:"picture_path"`
	ChafaBin    string `toml:"chafa"`
	Width       int    `toml:"width"`
}

func ParseFile(name string) (*Config, error) {
	err := createDefaultFileIfNotExists(name)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	conf := &Config{}

	err = toml.Unmarshal(data, conf)
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

		err = toml.NewEncoder(file).Encode(conf)
		if err != nil {
			return err
		}

		return file.Close()
	}
	return nil
}
