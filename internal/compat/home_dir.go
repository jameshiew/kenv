package compat

import (
	"os/user"
)

// UserHomeDir is equivalent to Go 1.12's ``os.UserHomeDir()``
func UserHomeDir() (string, error) {
	current, err := user.Current()
	if err != nil {
		return "", err
	}
	return current.HomeDir, nil
}
