package main

import (
	"dh/cmd"
	dh "dh/src"
	"fmt"
	"os"
)

func main() {
	dh.Config.Init()

	if err := cmd.RootCmd.Execute(); err != nil {
		if err != cmd.SilentErr {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
