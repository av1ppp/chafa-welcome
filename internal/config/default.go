package config

func getDefaultConfig() *Config {
	return &Config{
		ChafaBin: "chafa",
		Image: configImage{
			Source: "/home/user/Pictures/welcome.png",
			Size:   32,
		},
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
