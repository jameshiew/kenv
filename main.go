package main

import (
	"github.com/jameshiew/kenv/cmd"
)

var version string

func init() {
	if version == "" {
		version = "dev"
	}
}

func main() {
	cmd.Execute(version)
}
