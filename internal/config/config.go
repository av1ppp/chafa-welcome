package config

type Config struct {
	ChafaBin string       `toml:"chafa"`
	Offset   configOffset `toml:"offset"`
	Image    configImage  `toml:"image"`
	Body     configBody   `toml:"body"`
	Theme    configTheme  `toml:"theme"`
}

type configOffset struct {
	X int `toml:"x"`
	Y int `toml:"y"`
}

type configImage struct {
	Source string `toml:"source"`
	Size   int    `toml:"size"`
}

type configBody struct {
	Gap      int                `toml:"gap"`
	OS       configBodyOS       `toml:"os"`
	Kernel   configBodyKernel   `toml:"kernel"`
	Terminal configBodyTerminal `toml:"terminal"`
	Uptime   configBodyUptime   `toml:"uptime"`
	Packages configBodyPackages `toml:"packages"`
	Shell    configBodyShell    `toml:"shell"`
	CPU      configBodyCPU      `toml:"cpu"`
	Memory   configBodyMemory   `toml:"memory"`
	LocalIP  configBodyLocalIP  `toml:"local_ip"`
	GlobalIP configBodyGlobalIP `toml:"global_ip"`
}

type configBodyOS struct {
	Include bool `toml:"include"`
}

type configBodyKernel struct {
	Include bool `toml:"include"`
}

type configBodyTerminal struct {
	Include bool `toml:"include"`
}

type configBodyUptime struct {
	Include bool `toml:"include"`
}

type configBodyPackages struct {
	Include bool `toml:"include"`
}

type configBodyShell struct {
	Include bool `toml:"include"`
}

type configBodyCPU struct {
	Include bool `toml:"include"`
}

type configBodyMemory struct {
	Include bool `toml:"include"`
	Percent bool `toml:"percent"`
}

type configBodyLocalIP struct {
	Include bool `toml:"include"`
}

type configBodyGlobalIP struct {
	Include bool   `toml:"include"`
	Source  string `toml:"source"`
}

type configTheme struct {
	HeaderUsername  string `toml:"username"`
	HeaderAt        string `toml:"at"`
	HeaderHostname  string `toml:"hostname"`
	HeaderUnderline string `toml:"underline"`

	BodyKey       string `toml:"key"`
	BodySeparator string `toml:"separator"`
	BodyValue     string `toml:"value"`
}
