package sysinfo

import (
	"strings"
)

type SystemInfo struct {
	OS     string
	Kernel string
	Uptime string
}

func New() (*SystemInfo, error) {
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

	return &SystemInfo{
		OS:     os,
		Kernel: kernel,
		Uptime: uptime,
	}, nil
}

func (self *SystemInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString("OS: " + self.OS + "\n")
	builder.WriteString("Kernel: " + self.Kernel + "\n")
	builder.WriteString("Uptime: " + self.Uptime + "\n")
	// TODO remove last \n

	str := builder.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}
