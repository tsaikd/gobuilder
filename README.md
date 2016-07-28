# gobuilder
Go application builder

[![Build Status](https://travis-ci.org/tsaikd/gobuilder.svg?branch=master)](https://travis-ci.org/tsaikd/gobuilder)

## Why?

`go build` command works fine, but not enough. I need more information to embed
with application. e.g. **version**, **build time**, **revision**, **dependencies**

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

## Use gobuider with version constraint in build script
```
if ! gobuilder version -c ">=0.1" &>/dev/null ; then
	go get -u -v "github.com/tsaikd/gobuilder"
fi

gobuilder
```

## [Example](example) application output
```
$ gobuilder
$ ./example version -n
0.0.1
$ ./example version -c ">=1"
Error: current version "0.0.1" not in range ">=1"
$ ./example version
{
	"version": "0.0.1",
	"buildtime": "Fri, 29 Jul 2016 01:36:11 CST",
	"gitcommit": "0a470c",
	"godeps": {
		"Deps": [
			{
				"ImportPath": "github.com/spf13/cobra",
				"Rev": "f62e98d28ab7ad31d707ba837a966378465c7b57"
			},
			{
				"ImportPath": "github.com/spf13/pflag",
				"Rev": "1560c1005499d61b80f865c04d39ca7505bf7f0b"
			},
			{
				"ImportPath": "github.com/hashicorp/go-version",
				"Rev": "deeb027c13a95d56c7585df3fe29207208c6706e"
			},
			{
				"ImportPath": "github.com/spf13/viper",
				"Rev": "b53595fb56a492ecef90ee0457595a999eb6ec15"
			},
			{
				"ImportPath": "github.com/BurntSushi/toml",
				"Rev": "99064174e013895bbd9b025c31100bd1d9b590ca"
			},
			{
				"ImportPath": "github.com/fsnotify/fsnotify",
				"Rev": "a8a77c9133d2d6fd8334f3260d06f60e8d80a5fb"
			},
			{
				"ImportPath": "golang.org/x/sys",
				"Rev": "a646d33e2ee3172a661fc09bca23bb4889a41bc8"
			},
			{
				"ImportPath": "github.com/hashicorp/hcl",
				"Rev": "d8c773c4cba11b11539e3d45f93daeaa5dcf1fa1"
			},
			{
				"ImportPath": "github.com/magiconair/properties",
				"Rev": "b3f6dd549956e8a61ea4a686a1c02a33d5bdda4b"
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
			},
			{
				"ImportPath": "github.com/kardianos/osext",
				"Rev": "29ae4ffbc9a6fe9fb2bc5029050ce6996ea1d3bc"
			}
		],
		"GoVersion": "go1.6.3",
		"ImportPath": "github.com/tsaikd/gobuilder/example",
		"Rev": "0a470c2f69986eede7c59cbb671b786eca41d7e9"
	}
}
```
