package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"

	// load main entrypoint module
	_ "github.com/tsaikd/gobuilder/cmd"
	_ "github.com/tsaikd/gobuilder/cmd/dep"
)

func main() {
	cmder.Main()
}
