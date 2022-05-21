# SkyJet

[![Build Status](https://github.com/skyareas/skyjet/workflows/Release/badge.svg)](https://github.com/skyareas/skyjet/actions?query=workflow%3ARelease)
[![Go Reference](https://pkg.go.dev/badge/github.com/skyareas/skyjet.svg)](https://pkg.go.dev/github.com/skyareas/skyjet)

Skyjet is a batteries-included HTTP web framework for Golang.

> Skyjet is still in early development, and its API is subject to breaking changes, please don't use it in production.

## Contents

- [SkyJet](#skyJet)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [API Examples](#api-examples)
    - [Using Get, Post, Put, Patch, Delete and Options](#using-get-post-put-patch-delete-and-options)
    - [Using Router](#using-router)

## Installation

Skyjet can be installed as a Go module, to begin with, install Skyjet into your application:

```console
foo@bar$ go get -u github.com/skyareas/skyjet
```

then, import Skyjet in your application:

```go
import "github.com/skyareas/skyjet"
```

## Quick Start

Below is the hello world example using Skyjet:

```go
package main

import "github.com/skyareas/skyjet"

func main() {
    app := skyjet.SharedApp()
    app.Get("/", func(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
        return res.Send([]byte("Hello, Skyjet!"))
    })
    app.Run()
}
```

Build and run your application and open your browser at [localhost:8080](//localhost:8080)

## API Examples

A detailed documentation is not ready at the moment but the work is in progress!

### Using Get, Post, Put, Patch, Delete and Options

```go
import "github.com/skyareas/skyjet"

func main() {
    app := skyjet.SharedApp()
    app.Get("/path", getHandler)
    app.Post("/path", postHandler)
    app.Put("/path", putHandler)
    app.Delete("/path", deleteHandler)
    app.Patch("/path", patchHandler)
    app.Head("/path", headHandler)
    app.Options("/path", optionsHandler)
    app.Run()
}
```

### Using Router

```go
import "github.com/skyareas/skyjet"

func main() {
    app := skyjet.SharedApp()
	
    auth := skyjet.NewRouter()
    auth.Post("/login", loginHandler)
    app.Use("/auth", auth)

    blog := skyjet.NewRouter()
    blog.Post("/latest", latestArticlesHandler)
    app.Use("/blog", blog)

    app.Run()
}
```
