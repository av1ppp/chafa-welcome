package config

func getDefaultConfig() *Config {
	return &Config{
		ChafaBin: "chafa",
		Offset: configOffset{
			X: 1,
			Y: 0,
		},
		Image: configImage{
			Source: "/path/to/image.jpg",
			Size:   32,
		},
		Body: configBody{
			Gap:         2,
			AlignColumn: true,

			OS: configBodyOS{
				Include: true,
			},
			Kernel: configBodyKernel{
				Include: true,
			},
			Terminal: configBodyTerminal{
				Include: true,
			},
			Uptime: configBodyUptime{
				Include: true,
			},
			Packages: configBodyPackages{
				Include: true,
			},
			Shell: configBodyShell{
				Include: true,
			},
			CPU: configBodyCPU{
				Include: true,
			},
			Memory: configBodyMemory{
				Include: true,
				Percent: true,
			},
			LocalIP: configBodyLocalIP{
				Include: true,
			},
			GlobalIP: configBodyGlobalIP{
				Include: true,
				Source:  "ifconfig.me",
			},
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
