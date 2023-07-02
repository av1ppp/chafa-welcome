package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	ChafaBin string      `toml:"chafa"`
	Image    configImage `toml:"image"`
	Theme    configTheme `toml:"theme"`
}

type configImage struct {
	Source string `toml:"source"`
	Size   int    `toml:"size"`
}

type configTheme struct {
	HeaderUsername  string `toml:"header_username"`
	HeaderAt        string `toml:"header_at"`
	HeaderHostname  string `toml:"header_hostname"`
	HeaderUnderline string `toml:"header_underline"`

	BodyKey       string `toml:"body_key"`
	BodySeparator string `toml:"body_separator"`
	BodyValue     string `toml:"body_value"`
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

	conf := &Config{}

	err = toml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	err = validate(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func createDefaultFileIfNotExists(name string) error {
	_, err := os.Stat(name)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		conf := getDefaultConfig()
		file, err := os.Create(name)
		if err != nil {
			return err
		}

		enc := toml.NewEncoder(file)
		enc.Indent = ""
		err = enc.Encode(conf)
		if err != nil {
			return err
		}

		return file.Close()
	}
	return nil
}
