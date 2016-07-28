package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/cmd/build"
	"github.com/tsaikd/gobuilder/cmd/checkerror"
	"github.com/tsaikd/gobuilder/cmd/dep"
	"github.com/tsaikd/gobuilder/cmd/get"
	"github.com/tsaikd/gobuilder/cmd/restore"
)

// Module info
var Module = &cobrather.Module{
	Use:   "gobuilder",
	Short: "Go application builder, run action: restore -> get -> build",
	Commands: []*cobrather.Module{
		dep.Module,
		checkerror.Module,
		restore.Module,
		get.Module,
		build.Module,
		cobrather.VersionModule,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		runFuncs := []func(cmd *cobra.Command, args []string) error{}

		cmdModules := []*cobrather.Module{
			restore.Module,
			get.Module,
			build.Module,
		}
		depModules := cobrather.ListDeps(cobrather.OIncludeDepInCommand, cmdModules...)
		preRun := cobrather.GenRunE(depModules...)
		runFuncs = append(runFuncs, preRun)

		run := cobrather.GenRunE(cmdModules...)
		runFuncs = append(runFuncs, run)

		postRun := cobrather.GenPostRunE(cmdModules...)
		runFuncs = append(runFuncs, postRun)

		depPostRun := cobrather.GenPostRunE(depModules...)
		runFuncs = append(runFuncs, depPostRun)

		for _, fn := range runFuncs {
			if err := fn(cmd, args); err != nil {
				return err
			}
		}

		return nil
	},
}
