package main

import (
	plum "github.com/go-plum/plum"
)

func hello(ctx *plum.Context) {
	ctx.JSON(200, "hello from :"+ctx.Request.URL.String())
}
func main() {
	p := plum.New()
	p.GET("/hello", hello)

	r := p.Group("/1")
	r.GET("/hello", hello)

	r = p.Group("/2").Group("/4")
	r.GET("/hello", hello)

	r = p.Group("/2")
	r.GET("/hello", hello)

	rp := r.Group("/3")
	rp.GET("/hello", hello)

	p.Run(":8080")
}
