package main

import (
	"os"

	"github.com/fakturk/countbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
