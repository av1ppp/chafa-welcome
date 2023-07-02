package sysinfo

import (
	"github.com/klauspost/cpuid/v2"
)

func collectCPU() (string, error) {
	return cpuid.CPU.BrandName, nil
}
