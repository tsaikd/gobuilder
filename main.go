package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/gobuilder/cmd"
	"github.com/tsaikd/gobuilder/cmd/dep"
)

func main() {
	cmder.Main(
		*cmd.Module,
		*dep.Module,
	)
}
