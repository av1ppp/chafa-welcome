package config

import (
	"os"
)

type Config struct {
	picturePath string
	chafaBin    string
	chafaArgs   []string
}

func (self *Config) PicturePath() string {
	return self.picturePath
}

func (self *Config) ChafaBin() string {
	return self.chafaBin
}

func (self *Config) ChafaArgs() []string {
	return self.chafaArgs
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

	config, err := unmarshal(data)
	if err != nil {
		return nil, err
	}

	if err = validate(config); err != nil {
		return nil, err
	}

	return config, nil
}

func createDefaultFileIfNotExists(name string) error {
	_, err := os.Stat(name)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		return os.WriteFile(name, defaultConfigData, os.ModePerm)
	}
	return nil
}
