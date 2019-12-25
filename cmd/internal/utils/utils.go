package utils

import (
	"fmt"
	"os"
)

type errorPrinter interface {
	PrintErrf(format string, i ...interface{})
}

func RunAndExit(printer errorPrinter, fn func() error) {
	err := fn()
	if err != nil {
		printer.PrintErrf("Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func PrintAndExit(printer errorPrinter, fn func() (string, error)) {
	RunAndExit(printer, func() error {
		s, err := fn()
		if err != nil {
			return err
		}
		if s != "" { // avoid pointlessly printing a newline
			fmt.Println(s)
		}
		return nil
	})
}
