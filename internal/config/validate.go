package config

import (
	"os"

	"github.com/pkg/errors"
)

func validate(config *Config) error {
	// picturePath
	if config.PicturePath == "" {
		return errors.Wrap(&errorParameterMustBeSpecified{"Picture"}, validationErr.Error())
	}
	_, err := os.Stat(config.PicturePath)
	if err != nil {
		return errors.Wrap(err, validationErr.Error())
	}

	// chafaBin
	if config.ChafaBin == "" {
		return errors.Wrap(&errorParameterMustBeSpecified{"Chafa"}, validationErr.Error())
	}

	// width
	if config.Width == 0 {
		return errors.Wrap(&errorParameterMustBeSpecified{"Width"}, validationErr.Error())
	}

	return nil
}
