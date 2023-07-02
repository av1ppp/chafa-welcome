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
	var (
		info = SystemInfo{
			HeaderAt: "@",
			conf:     conf,
		}
		err error
	)

	if info.HeaderUsername, err = collectUsername(); err != nil {
		return nil, err
	}

	if info.HeaderHostname, err = collectHostname(); err != nil {
		return nil, err
	}

	info.HeaderUnderline = strings.Repeat("~", len(info.HeaderUsername)+1+len(info.HeaderHostname))

	if info.OS, err = collectIfInclude(conf.Body.OS.Include, conf, collectOS); err != nil {
		return nil, err
	}

	if info.Kernel, err = collectIfInclude(conf.Body.Kernel.Include, conf, collectKernel); err != nil {
		return nil, err
	}

	if info.Uptime, err = collectIfInclude(conf.Body.Uptime.Include, conf, collectUptime); err != nil {
		return nil, err
	}

	if info.Packages, err = collectIfInclude(conf.Body.Packages.Include, conf, collectPackages); err != nil {
		return nil, err
	}

	if info.Shell, err = collectIfInclude(conf.Body.Shell.Include, conf, collectShell); err != nil {
		return nil, err
	}

	if info.Terminal, err = collectIfInclude(conf.Body.Terminal.Include, conf, collectTerminal); err != nil {
		return nil, err
	}

	if info.CPU, err = collectIfInclude(conf.Body.CPU.Include, conf, collectCPU); err != nil {
		return nil, err
	}

	if info.Memory, err = collectIfInclude(conf.Body.Memory.Include, conf, collectMemory); err != nil {
		return nil, err
	}

	if info.LocalIP, err = collectIfInclude(conf.Body.LocalIP.Include, conf, collectLocalIP); err != nil {
		return nil, err
	}

	if info.GlobalIP, err = collectIfInclude(conf.Body.GlobalIP.Include, conf, collectGlobalIP); err != nil {
		return nil, err
	}

	return &info, nil
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

	body := [][2]string{
		{"OS", self.OS},
		{"Kernel", self.Kernel},
		{"Uptime", self.Uptime},
		{"Packages", self.Packages},
		{"Shell", self.Shell},
		{"Terminal", self.Terminal},
		{"CPU", self.CPU},
		{"Memory", self.Memory},
		{"Local IP", self.LocalIP},
		{"Global IP", self.GlobalIP},
	}

	for _, row := range body {
		if row[1] == "" {
			continue
		}
		builder.WriteString(colorKey.Sprint(row[0]))
		builder.WriteString(colorSeparator.Sprint(":") + " ")
		builder.WriteString(colorValue.Sprint(row[1]) + "\n")
	}

	str := builder.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}
