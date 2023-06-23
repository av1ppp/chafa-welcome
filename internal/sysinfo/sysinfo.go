package sysinfo

import (
	"strings"
)

type SystemInfo struct {
	Header   string
	OS       string
	Kernel   string
	Uptime   string
	Packages string
}

func New() (*SystemInfo, error) {
	header, err := collectHeader()
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

	// TODO add: username, hostname, local ip, shell, cpu, gpu, ram

	return &SystemInfo{
		Header:   header,
		OS:       os,
		Kernel:   kernel,
		Uptime:   uptime,
		Packages: packages,
	}, nil
}

func (self *SystemInfo) String() string {
	builder := strings.Builder{}

	builder.WriteString(self.Header + "\n")
	builder.WriteString(strings.Repeat("~", len(self.Header)) + "\n")

	builder.WriteString("OS: " + self.OS + "\n")
	builder.WriteString("Kernel: " + self.Kernel + "\n")
	builder.WriteString("Uptime: " + self.Uptime + "\n")
	builder.WriteString("Packages: " + self.Packages + "\n")
	// TODO remove last \n

	str := builder.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}
