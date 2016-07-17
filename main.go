package main

import (
	"os"

	"github.com/tsaikd/gobuilder/cmd"
	"github.com/tsaikd/gobuilder/cmd/flags"
)

func main() {
	rootCommand := cmd.Module.MustNewRootCommand(flags.Viper)
	rootCommand.SilenceUsage = true
	if err := rootCommand.Execute(); err != nil {
		os.Exit(-1)
	}
}
