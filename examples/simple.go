package main

import (
	"github.com/nazhard/hayane"
	//"github.com/nazhard/hayane/middleware"
)

func main() {
	app := hayane.Create()

	//app.Add(middleware.Logging)

	app.GET("/", func(ctx *hayane.Context) {
		ctx.String("Hello Mina-san!")
	})

	app.GET("/nazhard", func(ctx *hayane.Context) {
		ctx.JSON(map[string]string{
			"name":   "nazhard",
			"github": "https://github.com/nazhard",
		})
	})

	app.Run(":8080")
}
