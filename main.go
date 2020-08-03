package main

import (
	"os"

	"github.com/elastic/beats/libbeat/version"

	"github.com/mirerosystem/lsbeat/cmd"

	_ "github.com/mirerosystem/lsbeat/include"
)

var qualifier string

func main() {
	version.SetQualifier(qualifier)

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
