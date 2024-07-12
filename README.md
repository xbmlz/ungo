# ungo

[![Build Status](https://github.com/xbmlz/ungo/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/features/actions)
[![Coverage Status](https://coveralls.io/repos/github/xbmlz/ungo/badge.svg?branch=main)](https://coveralls.io/github/xbmlz/ungo?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/xbmlz/ungo)](https://goreportcard.com/report/github.com/xbmlz/ungo)
[![Go Doc](https://godoc.org/github.com/xbmlz/ungo?status.svg)](https://godoc.org/github.com/xbmlz/ungo)
[![Code Size](https://img.shields.io/github/languages/code-size/xbmlz/ungo.svg?style=flat-square)](https://github.com/xbmlz/ungo)
[![Release](https://img.shields.io/github/release/xbmlz/ungo.svg?style=flat-square)](https://github.com/xbmlz/ungo/releases)

Unlock the potential of your golang development journey with UnJS - where innovation meets simplicity, and possibilities become limitless.


## Usage

```go
package main

import (
	"flag"

	"github.com/xbmlz/ungo/cfg"
	"github.com/xbmlz/ungo/server"
)

var configFile = flag.String("c", "config.yaml", "config file path")

type Config struct {
	Server server.Config
}

func main() {
	flag.Parse()

	config := &Config{}
	cfg.MustLoad(*configFile, config)

	srv := server.MustNewHTTPServer(config.Server)
	srv.Run()
}


```