package config

import (
	"fmt"
	"strings"
)

var chafaArgs = []string{
	"Background", // --bg
	"Colors",     // --colors
}

func isChafaArg(arg string) bool {
	for _, arg_ := range chafaArgs {
		if arg == arg_ {
			return true
		}
	}
	return false
}

func isEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func splitRow(row string, rowN int) (key string, value string, err error) {
	rowSplitted := strings.SplitN(row, " ", 2)
	if len(rowSplitted) != 2 {
		return key, value, &errorIncorrectSyntax{rowN, row}
	}

	key = rowSplitted[0]
	value = rowSplitted[1]
	return
}

func setupKeyValueForConfig(key, value string, config *Config, row string, rowN int) error {
	if key == "Picture" {
		config.picturePath = value
		return nil
	}

	if key == "Chafa" {
		config.chafaBin = value
		return nil
	}

	if isChafaArg(key) {
		if len(value) > 0 {
			config.chafaArgs = append(config.chafaArgs, fmt.Sprintf("--%s=%s", key, value))
		}
		return nil
	}

	return &errorIncorrectParameter{
		key:   key,
		value: value,
		row:   row,
		rowN:  rowN,
	}
}

func unmarshal(data []byte) (*Config, error) {
	config := &Config{
		chafaArgs: []string{},
	}

	rows := strings.Split(string(data), "\n")

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if isEmptyString(row) {
			continue
		}

		key, value, err := splitRow(row, i+1)
		if err != nil {
			return nil, err
		}

		err = setupKeyValueForConfig(key, value, config, row, i+1)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
