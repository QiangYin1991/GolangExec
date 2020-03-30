package main

import (
	"fmt"
	"html/template"
	"net/http"
	"src/7days/Web/gee"
	"time"
)

type student struct {
	Name string
	Age  uint8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main()  {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFunMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	
	stu1 := &student{
		Name: "stu1",
		Age:  20,
	}
	stu2 := &student{
		Name: "stu2",
		Age:  30,
	}

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "students",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "CurrentTime",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}