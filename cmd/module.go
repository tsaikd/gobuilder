package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/gobuilder/cmd/modBuild"
	"github.com/tsaikd/gobuilder/cmd/modCheckError"
	"github.com/tsaikd/gobuilder/cmd/modCheckFmt"
	"github.com/tsaikd/gobuilder/cmd/modCheckRedundant"
	"github.com/tsaikd/gobuilder/cmd/modDep"
	"github.com/tsaikd/gobuilder/cmd/modFlags"
	"github.com/tsaikd/gobuilder/cmd/modGet"
	"github.com/tsaikd/gobuilder/cmd/modRestore"
	"github.com/tsaikd/gobuilder/cmd/modTest"
)

// command line flags
var (
	flagCheck = &cobrather.BoolFlag{
		Name:      "check",
		ShortHand: "c",
		Default:   false,
		Usage:     "Add check actions in building chain: checkfmt -> <BUILD ACTIONS> -> checkerror -> checkredundant",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "gobuilder",
	Short: "Go application builder, run action: restore -> get -> build -> [test]",
	Commands: []*cobrather.Module{
		modDep.Module,
		modCheckFmt.Module,
		modRestore.Module,
		modGet.Module,
		modBuild.Module,
		modTest.Module,
		modCheckError.Module,
		modCheckRedundant.Module,
		cobrather.VersionModule,
	},
	Dependencies: []*cobrather.Module{
		modFlags.Module,
	},
	Flags: []cobrather.Flag{
		flagCheck,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		runFuncs := []func(cmd *cobra.Command, args []string) error{}

		cmdModules := []*cobrather.Module{}
		if flagCheck.Bool() {
			cmdModules = append(cmdModules,
				modCheckFmt.Module,
			)
		}
		cmdModules = append(cmdModules,
			modRestore.Module,
			modGet.Module,
			modBuild.Module,
		)
		if modFlags.Test() {
			cmdModules = append(cmdModules,
				modTest.Module,
			)
		}
		if flagCheck.Bool() {
			cmdModules = append(cmdModules,
				modCheckError.Module,
				modCheckRedundant.Module,
			)
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
