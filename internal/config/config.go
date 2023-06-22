package config

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	picturePath string
}

func (self *Config) PicturePath() string {
	return self.picturePath
}

var incorrectFileSyntaxErr = errors.New("incorrect file syntax")
var validateErr = errors.New("failed to validate config")

func ParseFile(name string) (*Config, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	rows := strings.Split(string(data), "\n")
	for _, row := range rows {
		if len(strings.TrimSpace(row)) == 0 {
			continue
		}

		splitRow := strings.SplitN(row, "=", 2)
		if len(splitRow) != 2 {
			return nil, incorrectFileSyntaxErr
		}

		key := splitRow[0]
		value := splitRow[1]

		switch key {
		case "picture":
			config.picturePath = value
		default:
			return nil, incorrectFileSyntaxErr
		}
	}

	if err = validate(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validate(config *Config) error {
	_, err := os.Stat(config.picturePath)
	if err != nil {
		return errors.Join(validateErr, err)
	}

	return nil
}
