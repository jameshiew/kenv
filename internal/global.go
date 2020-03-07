package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func globalConf() (string, error) {
	root, err := RootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, configurationFilename), nil
}

func GlobalVersion() (string, error) {
	globalConfPath, err := globalConf()
	if err != nil {
		return "", err
	}
	version, err := ioutil.ReadFile(globalConfPath)
	if err != nil {
		if os.IsNotExist(err) {
			return systemVersion, nil
		}
		return "", err
	}
	return string(version), nil
}

func SetGlobalVersion(version string) error {
	exists, err := existsVersion(version)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("version %v is not installed", version)
	}
	globalConfPath, err := globalConf()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(globalConfPath, []byte(version), 0644)
}
