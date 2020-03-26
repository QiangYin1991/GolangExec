package main

import (
	"net/http"
	"src/7days/Web/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello/:name", func(context *gee.Context) {
		context.String(http.StatusOK, "hello %s, you're at %s\n", context.Param("name"), context.Path)
	})

	r.POST("/login", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
