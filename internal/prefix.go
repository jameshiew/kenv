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
	return prefixDir(version)
}

func ExistingPrefixDir(version string) (string, error) {
	exists, err := existsVersion(version)
	if err != nil {
		return "", nil
	}
	if !exists {
		return "", fmt.Errorf("version %v is not installed", version)
	}
	return prefixDir(version)
}

func prefixDir(version string) (string, error) {
	versions, err := versionsDir()
	if err != nil {
		return "", nil
	}
	if version == systemVersion {
		return "", fmt.Errorf("no prefix for system kubectl")
	}
	return filepath.Join(versions, version), nil
}
