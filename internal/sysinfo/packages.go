package sysinfo

import (
	"fmt"
	"strings"

	"github.com/av1ppp/chafa-welcome/internal/config"
)

func collectPackages(conf *config.Config) (string, error) {
	results := []string{}
	count := 0

	// pacman
	data, err := execute("pacman", "-Qq", "--color", "never")
	if err == nil {
		count = strings.Count(data, "\n")
		results = append(results, fmt.Sprintf("%d pacman", count))
	}

	// apt
	data, err = execute("apt-cache", "pkgnames")
	if err == nil {
		count = strings.Count(data, "\n")
		results = append(results, fmt.Sprintf("%d apt", count))
	}

	// dpkg
	data, err = execute("dpkg-query", "-f", ".\n", "-W")
	if err == nil {
		count = strings.Count(data, "\n")
		results = append(results, fmt.Sprintf("%d dpkg", count))
	}

	// flatpak
	data, err = execute("flatpak", "list")
	if err == nil {
		count = strings.Count(data, "\n")
		results = append(results, fmt.Sprintf("%d flatpak", count))
	}

	return strings.Join(results, ", "), nil
}
