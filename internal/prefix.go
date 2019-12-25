package internal

import (
	"fmt"
	"path/filepath"
)

func CurrentPrefixDir() (string, error) {
	version, err := CurrentVersion()
	if err != nil {
		return "", err
	}
	return PrefixDir(version)
}

func PrefixDir(version string) (string, error) {
	versions, err := versionsDir()
	if err != nil {
		return "", nil
	}
	if version == "system" {
		return "", fmt.Errorf("no prefix for system kubectl")
	}
	exists, err := existsVersion(version)
	if err != nil {
		return "", nil
	}
	if !exists {
		return "", fmt.Errorf("version %v is not installed", version)
	}
	return filepath.Join(versions, version), nil
}
