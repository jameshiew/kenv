package internal

import (
	"fmt"
	"os/exec"
)

func CurrentVersion() (string, error) {
	shell, err := ShellVersion()
	if err != nil {
		return "", err
	}
	if shell != "" {
		exists, err := existsVersion(shell)
		if err != nil {
			return "nil", err
		}
		if !exists {
			return "", fmt.Errorf("version %v is not installed", shell)
		}
		return shell, nil
	}
	local, err := LocalVersion()
	if err != nil {
		return "", err
	}
	if local != "" {
		exists, err := existsVersion(local)
		if err != nil {
			return "nil", err
		}
		if !exists {
			return "", fmt.Errorf("version %v is not installed", local)
		}
		return local, nil
	}
	return GlobalVersion()
}

func existsVersion(version string) (bool, error) {
	if version == "system" {
		cmd := exec.Command("command", "-v", "kubectl")
		err := cmd.Run()
		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return false, nil
			}
			return false, err
		}
		return true, nil
	}
	installed, err := InstalledVersions()
	if err != nil {
		return false, err
	}
	for _, v := range installed {
		if v == version {
			return true, nil
		}
	}
	return false, nil
}
