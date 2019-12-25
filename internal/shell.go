package internal

import (
	"fmt"
	"os"
)

const kenvVersion = "KENV_VERSION"

func ShellVersion() (string, error) {
	version := os.Getenv(kenvVersion)
	if version == "" {
		return "", nil
	}
	exists, err := existsVersion(version)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("version %v is not installed", version)
	}
	return version, nil
}
