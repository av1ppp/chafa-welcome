package sysinfo

import (
	"github.com/klauspost/cpuid/v2"

	"github.com/av1ppp/chafa-welcome/internal/config"
)

func collectCPU(conf *config.Config) (string, error) {
	return cpuid.CPU.BrandName, nil
}
