package main

import "github.com/nazhard/http"

func main() {
	r := http.New()

	r.GET("/", func(ctx http.Context) error {
		return ctx.String("hello")
	})
	r.GET("/nazhard", func(ctx http.Context) error {
		return ctx.String("his link: https://github.com/nazhard")
	})

	r.Serve(":8080")
}
