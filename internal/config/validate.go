package config

import (
	"os"

	"github.com/pkg/errors"
)

func validate(config *Config) error {
	// picturePath
	if config.Image.Source == "" {
		return errors.Wrap(&errorParameterMustBeSpecified{"image.source"}, validationErr.Error())
	}
	_, err := os.Stat(config.Image.Source)
	if err != nil {
		return errors.Wrap(err, validationErr.Error())
	}

	// chafaBin
	if config.ChafaBin == "" {
		return errors.Wrap(&errorParameterMustBeSpecified{"chafa"}, validationErr.Error())
	}

	// width
	if config.Image.Size == 0 {
		return errors.Wrap(&errorParameterMustBeSpecified{"image.size"}, validationErr.Error())
	}

	return nil
}
