Translations: [简体中文](README_zh.md)  | [English](README.md)


# Plum 

Plum 是一个简单，通用，高效，稳定，高性能，可扩展，无其他依赖库的 golang HTTP 库 .
这是一个标准库的封装.
 
## 为什么？
golang 标准库非常强大。快速，易用，且有一个非常好的 API.

尤其是在 1.22 之中增强了 HTTP 路由功能，plum 目标是高度兼容 net/http，增加处理能力。


## Goals
Implement a high performance, scalable, and stable http library with a simple, generic idea in a pure standard library way.
Plum is also an ideal one learning warehouse

## Principles

    - 零依赖，只使用 golang 标准库。
    - Simple 简单
    - Genneral 通用
    - Stable 稳定
    - Expandability 可扩展
	

## Features

- 零依赖，只使用 golang 标准库。
- 结构化日志接口，默认实现[log/slog].
- 使用最新的标准库路由 http.ServeMux.
- 可定制：middlewares,logger ,binding,render 等等。

# Getting Started

## Installation

You need Go version 1.22 or higher.

```shell
go get -u github.com/go-plum/plum
```

## Usage

主要使用上类似于 gin，且兼容部分或者全部 gin 的中间件，日志，binding，render .

 
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

	p.Run(":8080") // go p.Run(":8080") ... 
}


```

### 参考仓库 
+ [gin](https://github.com/gin-gonic/gin)
+ [kratos](https://github.com/go-kratos/kratos)