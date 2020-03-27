package main

import (
	"log"
	"net/http"
	"src/7days/Web/gee"
	"time"
)

func only4V2() gee.HandleFunc {
	return func(context *gee.Context) {
		t := time.Now()
		context.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(context *gee.Context) {
			context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(only4V2())
	{
		v2.GET("/hello/:name", func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s, you're at %s\n", context.Param("name"), context.Path)
		})

		v2.POST("/login", func(context *gee.Context) {
			context.JSON(http.StatusOK, gee.H{
				"username": context.PostForm("username"),
				"password": context.PostForm("password"),
			})
		})

		v2.GET("/assets/*filepath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
		})
	}

	r.Run(":9999")
}
