package internal

import (
	"io/ioutil"
)

func InstalledVersions() ([]string, error) {
	versionsPath, err := versionsDir()
	if err != nil {
		return nil, err
	}
	fileInfos, err := ioutil.ReadDir(versionsPath)
	if err != nil {
		return nil, err
	}
	var versions []string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			versions = append(versions, fileInfo.Name())
		}
	}
	return versions, nil
}
