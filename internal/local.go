package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const configurationFilename = ".kenv.yml"

type configuration struct {
	Version string `yaml:"version"`
}

const filesystemRoot = string(os.PathSeparator)

func LocalVersion() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for ; current != filesystemRoot; current = filepath.Dir(current) {
		contents, err := ioutil.ReadFile(filepath.Join(current, configurationFilename))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return "", err
		}
		var local configuration
		err = yaml.Unmarshal(contents, &local)
		if err != nil {
			return "", err
		}
		return local.Version, nil
	}
	// reached filesystemRoot without finding a configuration file
	return "", nil
}

func SetLocalVersion(version string) error {
	current, err := os.Getwd()
	if err != nil {
		return err
	}
	if current == filesystemRoot {
		return fmt.Errorf("may not set a local kubectl version for the root directory")
	}

	exists, err := existsVersion(version)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("version %v is not installed", version)
	}
	yml, err := yaml.Marshal(configuration{Version: version})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configurationFilename, yml, 0644)
}
