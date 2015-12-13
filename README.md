# gobuilder
Go application builder

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
	"buildtime": "Mon, 14 Dec 2015 06:34:08 CST",
	"gitcommit": "2afd48",
	"godeps": {
		"Deps": [
			{
				"Comment": "v0.8.7-49-gcdaedc6",
				"ImportPath": "github.com/Sirupsen/logrus",
				"Rev": "cdaedc68f2894175ac2b3221869685602c759e71"
			},
			{
				"Comment": "1.2.0",
				"ImportPath": "github.com/codegangsta/cli",
				"Rev": "565493f259bf868adb54d45d5f4c68d405117adf"
			},
			{
				"Comment": "v1.0-rc1-10-g33e0aa1",
				"ImportPath": "github.com/codegangsta/inject",
				"Rev": "33e0aa1cb7c019ccc3fbe049a8262a6403d30504"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib/cliutil/cmdutil",
				"Rev": "b0111ac77cfc517da0e1942938966e4f0a699889"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib/cliutil/flagutil",
				"Rev": "b0111ac77cfc517da0e1942938966e4f0a699889"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib/errutil",
				"Rev": "b0111ac77cfc517da0e1942938966e4f0a699889"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib/logrusutil",
				"Rev": "b0111ac77cfc517da0e1942938966e4f0a699889"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib/logutil",
				"Rev": "b0111ac77cfc517da0e1942938966e4f0a699889"
			},
			{
				"ImportPath": "github.com/tsaikd/KDGoLib/version",
				"Rev": "b0111ac77cfc517da0e1942938966e4f0a699889"
			}
		],
		"GoVersion": "go1.5.1",
		"ImportPath": "github.com/tsaikd/gobuilder/example"
	}
}
```
