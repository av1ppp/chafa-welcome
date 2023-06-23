package config

import (
	"os"

	"github.com/pkg/errors"
)

func validate(config *Config) error {
	// picturePath
	if config.picturePath == "" {
		return errors.Wrap(&errorParameterMustBeSpecified{"Picture"}, validationErr.Error())
	}
	_, err := os.Stat(config.picturePath)
	if err != nil {
		return errors.Wrap(err, validationErr.Error())
	}

	// chafaBin
	if config.chafaBin == "" {
		return errors.Wrap(&errorParameterMustBeSpecified{"Chafa"}, validationErr.Error())
	}

	// width
	if config.width == 0 {
		return errors.Wrap(&errorParameterMustBeSpecified{"Width"}, validationErr.Error())
	}

	return nil
}
