package config

func getDefaultConfig() *Config {
	return &Config{
		PicturePath: "/home/user/Pictures/welcome.png",
		ChafaBin:    "chafa",
		Width:       32,
	}
}
