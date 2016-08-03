package main

import (
	"os"

	"github.com/tsaikd/gobuilder/cmd"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
)

func main() {
	rootCommand := cmd.Module.MustNewRootCommand(modFlags.Viper)
	rootCommand.SilenceUsage = true
	if err := rootCommand.Execute(); err != nil {
		os.Exit(-1)
	}
}
