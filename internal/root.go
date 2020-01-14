package internal

import (
	"path/filepath"

	"github.com/jameshiew/kenv/internal/compat"
)

const (
	kubeDir = ".kube"
	kenvDir = "kenv"
)

func RootDir() (string, error) {
	home, err := compat.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, kubeDir, kenvDir), nil
}

func versionsDir() (string, error) {
	root, err := RootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "versions"), nil
}

func shimsDir() (string, error) {
	root, err := RootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "shims"), nil
}
