<img src="assets\taxediologolandscape.jpg" alt="drawing" width="200"/>

<h1 align="center">
  Tiologger
</h1>

<h3 align="center">
  <a href="https://taxed.io">taxed.io</a>
</h3>

![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/taxedio/tiologger/main?style=for-the-badge) ![GitHub](https://img.shields.io/github/license/taxedio/iso3166?style=for-the-badge) ![GitLab Release (custom instance)](https://img.shields.io/gitlab/v/release/taxedio/tiologger?include_prereleases&style=for-the-badge) ![Gitlab pipeline status](https://img.shields.io/gitlab/pipeline-status/taxedio/tiologger?branch=main&style=for-the-badge) ![Gitlab code coverage](https://img.shields.io/gitlab/coverage/taxedio/tiologger/main?style=for-the-badge) ![GitHub contributors](https://img.shields.io/github/contributors/taxedio/tiologger?style=for-the-badge)

[![Go Report Card](https://goreportcard.com/badge/github.com/taxedio/tiologger)](https://goreportcard.com/report/github.com/taxedio/tiologger)

# Introduction

This package creates a basic logging service.

# Basic Example

```GO
package app

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/taxedio/tiologger"
)

var (
	Router = gin.Default()
	EnvLoc = ".env"
)

func StartApplication() {
	Router = MapUrls(Router)
	logger.Info("about to start the application...")
	if err := Router.Run(); err != nil {
		logger.Critical("application.go error running router", err)
		panic(err)
	}
}
```

**console**:

```stdout
{"level":"info","time":"2022-04-07T12:12:38.079+0100","msg":"about to start the application..."}
```
