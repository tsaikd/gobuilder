# gobuilder
Go application builder

[![Build Status](https://travis-ci.org/tsaikd/gobuilder.svg?branch=master)](https://travis-ci.org/tsaikd/gobuilder)

## Install
```
go get -v "github.com/tsaikd/gobuilder"
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
	"buildtime": "Thu, 07 Apr 2016 23:11:53 CST",
	"gitcommit": "e8579c",
	"godeps": {
		"Deps": [
			{
				"ImportPath": "github.com/codegangsta/cli",
				"Rev": "565493f259bf868adb54d45d5f4c68d405117adf"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib",
				"Rev": "eabe0bda1bd0c304b889d5b34ad1e99f7808d93f"
			},
			{
				"ImportPath": "github.com/Sirupsen/logrus",
				"Rev": "4b6ea7319e214d98c938f12692336f7ca9348d6b"
			},
			{
				"ImportPath": "github.com/codegangsta/inject",
				"Rev": "33e0aa1cb7c019ccc3fbe049a8262a6403d30504"
			}
		],
		"GoVersion": "go1.6",
		"ImportPath": "github.com/tsaikd/gobuilder/example",
		"Rev": "e8579c0205c585d33151a203235ea65515daca4b"
	}
}
```
