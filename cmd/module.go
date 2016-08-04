package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/gobuilder/cmd/modBuild"
	"github.com/tsaikd/gobuilder/cmd/modCheckError"
	"github.com/tsaikd/gobuilder/cmd/modCheckFmt"
	"github.com/tsaikd/gobuilder/cmd/modDep"
	"github.com/tsaikd/gobuilder/cmd/modGet"
	"github.com/tsaikd/gobuilder/cmd/modRestore"
	"github.com/tsaikd/gobuilder/deputil"
)

// command line flags
var (
	flagCheck = &cobrather.BoolFlag{
		Name:      "check",
		ShortHand: "c",
		Default:   false,
		Usage:     "Run check actions before build actions: dep -> checkerror -> checkfmt",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "gobuilder",
	Short: "Go application builder, run action: restore -> get -> build",
	Commands: []*cobrather.Module{
		modDep.Module,
		modCheckError.Module,
		modCheckFmt.Module,
		modRestore.Module,
		modGet.Module,
		modBuild.Module,
		cobrather.VersionModule,
	},
	Flags: []cobrather.Flag{
		flagCheck,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		runFuncs := []func(cmd *cobra.Command, args []string) error{}

		if err := deputil.Check("", false); err != nil {
			if deputil.ErrorDepRevMismatch4.In(err) {
				return errutil.New("Check dependencies failed")
			}
			return err
		}

		cmdModules := []*cobrather.Module{}
		if flagCheck.Bool() {
			cmdModules = append(cmdModules,
				modCheckError.Module,
				modCheckFmt.Module,
			)
		}
		cmdModules = append(cmdModules,
			modRestore.Module,
			modGet.Module,
			modBuild.Module,
		)

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
