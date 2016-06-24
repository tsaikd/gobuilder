# gobuilder
Go application builder

[![Build Status](https://travis-ci.org/tsaikd/gobuilder.svg?branch=master)](https://travis-ci.org/tsaikd/gobuilder)

## Install
```
go get -u -v "github.com/tsaikd/gobuilder"
```

## Insert [version package](https://github.com/tsaikd/KDGoLib/tree/master/version) code

See [example](example) for usage

## Use gobuider to compile your application
```
gobuilder
```

## Example output
```
$ gobuilder
$ ./example version
{
	"version": "0.0.1",
	"buildtime": "Mon, 13 Jun 2016 02:48:34 CST",
	"gitcommit": "b0ab42",
	"godeps": {
		"Deps": [
			{
				"ImportPath": "github.com/tsaikd/KDGoLib",
				"Rev": "90b33fbfd7bf82557a32a829e7286df3bd25d277"
			},
			{
				"ImportPath": "gopkg.in/urfave/cli.v2",
				"Rev": "1b5ad735df034545a2ce018d348a814d406fc258"
			}
		],
		"GoVersion": "go1.6.2",
		"ImportPath": "github.com/tsaikd/gobuilder/example",
		"Rev": "b0ab42c30a95bcc4e094ac3cc040c4911b3da6db"
	}
}
```
