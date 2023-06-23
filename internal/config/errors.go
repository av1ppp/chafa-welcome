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

type errorIncorrectParameter struct {
	key, value string
	rowN       int
	row        string
}

func (self *errorIncorrectParameter) Error() string {
	return fmt.Sprintf("incorrect parameter (key='%s' value='%s' row=%d row-value='%s')", self.key, self.value, self.rowN, self.row)
}

type errorIncorrectSyntax struct {
	rowN int
	row  string
}

func (self *errorIncorrectSyntax) Error() string {
	return fmt.Sprintf("incorrect config syntax (row=%d row-value='%s')", self.rowN, self.row)
}

var validationErr = errors.New("failed to validate config")
