package config

func getDefaultConfig() *Config {
	return &Config{
		PicturePath: "/home/user/Pictures/welcome.png",
		ChafaBin:    "chafa",
		Width:       32,
		Theme: configTheme{
			HeaderUsername:  "fgred bold",
			HeaderAt:        "bold",
			HeaderHostname:  "fgred bold",
			HeaderUnderline: "",

			BodyKey:       "fgblue bold",
			BodySeparator: "",
			BodyValue:     "",
		},
	}
}
