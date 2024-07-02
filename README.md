
Translations: [English](README.md) | [简体中文](README_zh.md)

# Plum 

Plum is a Simple,General,Highly efficient,Stable,High-performance,Expandability, zero-dependency HTTP library for Go.

It's a thin wrapper around the standard library .

## Why?
The Go standard library is very powerful. Fast, easy to use, and with a good API.

In particular, with the improved HTTP routing in 1.22, plum aims to maintain net/http compatibility and enhance processing power


## Goals
Implement a high performance, scalable, and stable http library with a simple, generic idea in a pure standard library way.
Plum is also an ideal one learning warehouse

## Principles

	Zero-dependency
	Simple
	Genneral
	Stable
	Expandability
	


## Features

- Zero-dependency, only use the standard library.
- Structured logging interface,  default implementation[log/slog].
- routes Using new, improved http.ServeMux.
- customizable: middlewares,logger ,binding,render and so on.

# Getting Started

## Installation

You need Go version 1.22 or higher.

```shell
go get -u github.com/go-plum/plum
```

## Usage

It mainly uses middleware, logging, binding, render that is similar to gin and compatible with some or all of gin.


```go
package main

import (
	"fmt"

	"github.com/go-plum/plum"
)

func hello(ctx *plum.Context) error {
	fmt.Println("hello", ctx.Request.URL)
	ctx.JSON(200, "hello from :"+ctx.Request.URL.String())
	return nil
}
func main() {
	p := plum.New()
	p.GET("/hello", hello)

	r := p.Group("/1")
	r.GET("/hello", hello)

    r = p.Group("/2")
	r.GET("/hello", hello)

	r = p.Group("/2").Group("/3")
	r.GET("/hello", hello)

	p.Run(":8080") // go p.Run(":8080")
}


```
 
### Reference  
+ [gin](https://github.com/gin-gonic/gin)
+ [kratos](https://github.com/go-kratos/kratos)