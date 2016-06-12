package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/cmd"
	"github.com/tsaikd/gobuilder/cmd/build"
	"github.com/tsaikd/gobuilder/cmd/dep"
	"github.com/tsaikd/gobuilder/cmd/get"
	"github.com/tsaikd/gobuilder/cmd/restore"
)

func main() {
	cmder.Main(
		*cmd.Module,
		*dep.Module,
		*restore.Module,
		*get.Module,
		*build.Module,
	)
}
