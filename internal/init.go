package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Initialize() error {
	err := initializeKenvDirectory()
	if err != nil {
		return err
	}
	return nil
}

func InitializationHelpString() string {
	return `# Copy and paste the below into your .bashrc to have kenv initialize at shell start up
eval "$(kenv init -)"`
}

func ShellInitialization() (string, error) {
	shims, err := shimsDir()
	if err != nil {
		return "", err
	}
	// TODO: make shims derive from an envvar e.g. $KENV_ROOT, not hardcoded to HOME path
	return `export PATH="` + shims + `:${PATH}"

kenv() {
  if [[ $1 == 'shell' && -n $2 ]]
  then
    if [[ $2 != 'system' && -z "$(command kenv versions | grep $2)" ]]
    then
      echo "Error: version $2 is not installed"
      return
    fi
    export KENV_VERSION=$2
  else
    command kenv $@
  fi
}
`, nil
}

func initializeKenvDirectory() error {
	root, err := RootDir()
	if err != nil {
		return err
	}
	for _, dir := range []string{"versions", "shims"} {
		err := os.MkdirAll(filepath.Join(root, dir), 0o744)
		if err != nil {
			return err
		}
	}
	err = initializeShims()
	if err != nil {
		return err
	}
	return nil
}

func initializeShims() error {
	shims, err := shimsDir()
	if err != nil {
		return err
	}
	script := `#!/usr/bin/env bash

set -euo pipefail

if [[ $(kenv version) == "` + systemVersion + `" ]]
then
  SYSTEM_KUBECTL="$(which -a kubectl | sed -n 2p)"
	if [[ -z ${SYSTEM_KUBECTL} ]]
	then
	  echo "No version of kubectl is installed."
	  exit 1
	fi
  ${SYSTEM_KUBECTL} $@
else
  PREFIX="$(kenv prefix)"
  "${PREFIX}"/kubectl $@
fi
`
	err = ioutil.WriteFile(filepath.Join(shims, "kubectl"), []byte(script), 0o744)
	if err != nil {
		return err
	}
	return nil
}
