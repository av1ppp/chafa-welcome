package config

import (
	"os"
)

type Config struct {
	picturePath string
	chafaBin    string
	width       int
}

func (self *Config) PicturePath() string {
	return self.picturePath
}

func (self *Config) ChafaBin() string {
	return self.chafaBin
}

func (self *Config) Width() int {
	return self.width
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
