package config

import (
	"os"

	"github.com/pkg/errors"
)

func validate(config *Config) error {
	if config.picturePath != "" {
		_, err := os.Stat(config.picturePath)
		if err != nil {
			return errors.Wrap(err, validationErr.Error())
		}
	}

	if config.chafaBin == "" {
		return errors.Wrap(&errorParameterMustBeSpecefied{"Chafa"}, validationErr.Error())
	}

	return nil
}
