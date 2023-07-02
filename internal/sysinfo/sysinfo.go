package sysinfo

import (
	"github.com/av1ppp/chafa-welcome/internal/config"
	"strings"
)

type SystemInfo struct {
	HeaderUsername  string
	HeaderAt        string
	HeaderHostname  string
	HeaderUnderline string

	OS       string
	Kernel   string
	Uptime   string
	Packages string
	Shell    string
	Terminal string
	CPU      string
	Memory   string
	LocalIP  string
	GlobalIP string

	conf *config.Config
}

func Collect(conf *config.Config) (*SystemInfo, error) {
	username, err := collectUsername()
	if err != nil {
		return nil, err
	}

	hostname, err := collectHostname()
	if err != nil {
		return nil, err
	}

	os, err := collectOS()
	if err != nil {
		return nil, err
	}

	kernel, err := collectKernel()
	if err != nil {
		return nil, err
	}

	uptime, err := collectUptime()
	if err != nil {
		return nil, err
	}

	packages, err := collectPackages()
	if err != nil {
		return nil, err
	}

	shell, err := collectShell()
	if err != nil {
		return nil, err
	}

	terminal, err := collectTerminal()
	if err != nil {
		return nil, err
	}

	cpu, err := collectCPU()
	if err != nil {
		return nil, err
	}

	memory, err := collectMemory()
	if err != nil {
		return nil, err
	}

	localIP, err := collectLocalIP()
	if err != nil {
		return nil, err
	}

	globalIP, err := collectGlobalIP()
	if err != nil {
		return nil, err
	}

	return &SystemInfo{
		HeaderUsername:  username,
		HeaderAt:        "@",
		HeaderHostname:  hostname,
		HeaderUnderline: strings.Repeat("~", len(username)+1+len(hostname)),

		OS:       os,
		Kernel:   kernel,
		Uptime:   uptime,
		Packages: packages,
		Shell:    shell,
		Terminal: terminal,
		CPU:      cpu,
		Memory:   memory,
		GlobalIP: globalIP,
		LocalIP:  localIP,

		conf: conf,
	}, nil
}

func (self *SystemInfo) String() string {
	colorUsername := themeToColor(self.conf.Theme.HeaderUsername)
	colorAt := themeToColor(self.conf.Theme.HeaderAt)
	colorHostname := themeToColor(self.conf.Theme.HeaderHostname)
	colorUnderline := themeToColor(self.conf.Theme.HeaderUnderline)

	colorKey := themeToColor(self.conf.Theme.BodyKey)
	colorSeparator := themeToColor(self.conf.Theme.BodySeparator)
	colorValue := themeToColor(self.conf.Theme.BodyValue)

	builder := strings.Builder{}

	builder.WriteString(colorUsername.Sprint(self.HeaderUsername))
	builder.WriteString(colorAt.Sprint(self.HeaderAt))
	builder.WriteString(colorHostname.Sprint(self.HeaderHostname) + "\n")
	builder.WriteString(colorUnderline.Sprint(self.HeaderUnderline) + "\n")

	body := map[string]string{
		"OS":        self.OS,
		"Kernel":    self.Kernel,
		"Uptime":    self.Uptime,
		"Packages":  self.Packages,
		"Shell":     self.Shell,
		"Terminal":  self.Terminal,
		"CPU":       self.CPU,
		"Memory":    self.Memory,
		"Local IP":  self.LocalIP,
		"Global IP": self.GlobalIP,
	}

	for key, value := range body {
		builder.WriteString(colorKey.Sprint(key))
		builder.WriteString(colorSeparator.Sprint(":") + " ")
		builder.WriteString(colorValue.Sprint(value) + "\n")
	}

	str := builder.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}
