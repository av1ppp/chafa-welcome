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
	Shell    string
	Terminal string
	CPU      string
	Memory   string
	LocalIP  string
	GlobalIP string
}

func Collect() (*SystemInfo, error) {
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

	// TODO add: username, hostname, local ip, shell, cpu, gpu, ram

	return &SystemInfo{
		Header:   header,
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
	builder.WriteString("Shell: " + self.Shell + "\n")
	builder.WriteString("Terminal: " + self.Terminal + "\n")
	builder.WriteString("CPU: " + self.CPU + "\n")
	builder.WriteString("Memory: " + self.Memory + "\n")
	builder.WriteString("Local IP: " + self.LocalIP + "\n")
	builder.WriteString("Global IP: " + self.GlobalIP + "\n")

	str := builder.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}
