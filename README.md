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
	"buildtime": "Sun, 17 Jul 2016 21:10:19 CST",
	"gitcommit": "25904f",
	"godeps": {
		"Deps": [
			{
				"ImportPath": "github.com/spf13/cobra",
				"Rev": "b24564e919247d7c870fe0ed3738c98d8741aca4"
			},
			{
				"ImportPath": "github.com/spf13/pflag",
				"Rev": "367864438f1b1a3c7db4da06a2f55b144e6784e0"
			},
			{
				"ImportPath": "github.com/spf13/viper",
				"Rev": "c1ccc378a054ea8d4e38d8c67f6938d4760b53dd"
			},
			{
				"ImportPath": "github.com/BurntSushi/toml",
				"Rev": "bec2dacf4b590d26237cfebff4471e21ce543494"
			},
			{
				"ImportPath": "github.com/fsnotify/fsnotify",
				"Rev": "a8a77c9133d2d6fd8334f3260d06f60e8d80a5fb"
			},
			{
				"ImportPath": "golang.org/x/sys",
				"Rev": "b518c298ac9dc94b6ac0757394f50d10c5dfa25a"
			},
			{
				"ImportPath": "github.com/hashicorp/hcl",
				"Rev": "d8c773c4cba11b11539e3d45f93daeaa5dcf1fa1"
			},
			{
				"ImportPath": "github.com/magiconair/properties",
				"Rev": "e2f061ecfdaca9f35b2e2c12346ffc526f138137"
			},
			{
				"ImportPath": "github.com/mitchellh/mapstructure",
				"Rev": "21a35fb16463dfb7c8eee579c65d995d95e64d1e"
			},
			{
				"ImportPath": "github.com/spf13/cast",
				"Rev": "27b586b42e29bec072fe7379259cc719e1289da6"
			},
			{
				"ImportPath": "github.com/spf13/jwalterweatherman",
				"Rev": "33c24e77fb80341fe7130ee7c594256ff08ccc46"
			},
			{
				"ImportPath": "gopkg.in/yaml.v2",
				"Rev": "e4d366fc3c7938e2958e662b4258c7a89e1f0e3e"
			}
		],
		"GoVersion": "go1.6.2",
		"ImportPath": "github.com/tsaikd/gobuilder/example",
		"Rev": "25904f7cbebbbcfa7ead35847907632c74baec9b"
	}
}
```
