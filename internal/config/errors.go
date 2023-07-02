package config

import (
	"errors"
	"fmt"
)

type errorParameterMustBeSpecified struct {
	param string
}

func (self *errorParameterMustBeSpecified) Error() string {
	return fmt.Sprintf("parameter '%s' must be specified", self.param)
}

var validationErr = errors.New("validation error")
