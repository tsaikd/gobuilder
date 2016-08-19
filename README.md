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
gobuilder version -c ">=0.1" &>/dev/null || go get -u -v "github.com/tsaikd/gobuilder"
gobuilder --check
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
	"goversion": "go1.7",
	"buildtime": "Fri, 19 Aug 2016 14:44:20 CST",
	"gitcommit": "b2b592",
	"godeps": {
		"Deps": [
			{
				"ImportPath": "github.com/spf13/cobra",
				"Rev": "7c674d9e72017ed25f6d2b5e497a1368086b6a6f",
				"RevTime": "2016-08-02 18:37:37 -0400"
			},
			{
				"ImportPath": "github.com/spf13/pflag",
				"Rev": "4f9190456aed1c2113ca51ea9b89219747458dc1",
				"RevTime": "2016-08-16 14:05:11 -0400"
			},
			{
				"ImportPath": "github.com/hashicorp/go-version",
				"Rev": "deeb027c13a95d56c7585df3fe29207208c6706e",
				"RevTime": "2016-07-25 14:20:58 -0700"
			},
			{
				"ImportPath": "github.com/spf13/viper",
				"Rev": "654fc7bb54d0c138ef80405ff577391f79c0c32d",
				"RevTime": "2016-08-16 10:09:34 +0200"
			},
			{
				"ImportPath": "github.com/BurntSushi/toml",
				"Rev": "99064174e013895bbd9b025c31100bd1d9b590ca",
				"RevTime": "2016-07-17 11:07:09 -0400"
			},
			{
				"ImportPath": "github.com/fsnotify/fsnotify",
				"Rev": "f12c6236fe7b5cf6bcf30e5935d08cb079d78334",
				"RevTime": "2016-08-15 23:15:41 -0600"
			},
			{
				"ImportPath": "golang.org/x/sys",
				"Rev": "a646d33e2ee3172a661fc09bca23bb4889a41bc8",
				"RevTime": "2016-07-17 07:19:31 +0000"
			},
			{
				"ImportPath": "github.com/hashicorp/hcl",
				"Rev": "d8c773c4cba11b11539e3d45f93daeaa5dcf1fa1",
				"RevTime": "2016-07-11 17:17:52 -0600"
			},
			{
				"ImportPath": "github.com/magiconair/properties",
				"Rev": "61b492c03cf472e0c6419be5899b8e0dc28b1b88",
				"RevTime": "2016-08-16 10:55:11 +0200"
			},
			{
				"ImportPath": "github.com/mitchellh/mapstructure",
				"Rev": "ca63d7c062ee3c9f34db231e352b60012b4fd0c1",
				"RevTime": "2016-08-08 11:12:53 -0700"
			},
			{
				"ImportPath": "github.com/spf13/afero",
				"Rev": "b28a7effac979219c2a2ed6205a4d70e4b1bcd02",
				"RevTime": "2016-08-16 10:07:57 +0200"
			},
			{
				"ImportPath": "github.com/pkg/sftp",
				"Rev": "a71e8f580e3b622ebff585309160b1cc549ef4d2",
				"RevTime": "2016-07-22 09:14:53 +1000"
			},
			{
				"ImportPath": "github.com/kr/fs",
				"Rev": "2788f0dbd16903de03cb8186e5c7d97b69ad387b",
				"RevTime": "2013-11-10 17:25:53 -0800"
			},
			{
				"ImportPath": "github.com/pkg/errors",
				"Rev": "a22138067af1c4942683050411a841ade67fe1eb",
				"RevTime": "2016-08-08 15:55:40 +1000"
			},
			{
				"ImportPath": "golang.org/x/crypto",
				"Rev": "9fbab14f903f89e23047b5971369b86380230e56",
				"RevTime": "2016-08-17 14:31:42 +0000"
			},
			{
				"ImportPath": "golang.org/x/text",
				"Rev": "d69c40b4be55797923cec7457fac7a244d91a9b6",
				"RevTime": "2016-08-16 09:21:53 +0000"
			},
			{
				"ImportPath": "github.com/spf13/cast",
				"Rev": "e31f36ffc91a2ba9ddb72a4b6a607ff9b3d3cb63",
				"RevTime": "2016-07-30 11:20:37 +0200"
			},
			{
				"ImportPath": "github.com/spf13/jwalterweatherman",
				"Rev": "33c24e77fb80341fe7130ee7c594256ff08ccc46",
				"RevTime": "2016-03-11 10:36:46 +0100"
			},
			{
				"ImportPath": "gopkg.in/yaml.v2",
				"Rev": "e4d366fc3c7938e2958e662b4258c7a89e1f0e3e",
				"RevTime": "2016-07-15 00:37:55 -0300"
			},
			{
				"ImportPath": "github.com/kardianos/osext",
				"Rev": "c2c54e542fb797ad986b31721e1baedf214ca413",
				"RevTime": "2016-08-10 17:15:26 -0700"
			}
		],
		"ImportPath": "github.com/tsaikd/gobuilder/example",
		"Rev": "b2b5920be5f242c9ab36e8004403f934e3de64a1",
		"RevTime": "2016-08-19 14:43:34 +0800"
	}
}
```
