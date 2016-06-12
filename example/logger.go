package main

import (
	"github.com/tsaikd/KDGoLib/logutil"
	"gopkg.in/urfave/cli.v2"
)

var (
	logger = logutil.DefaultLogger
)

func actionWrapper(action func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		if err := action(context); err != nil {
			logger.Fatalln(err)
		}
	}
}
