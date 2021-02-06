package main

import (
	"html/template"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// new template engine
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Funcs:        template.FuncMap{},
		DisableCache: true,
	})

	// kita pasang middleware
	r.Use(authMiddleware)

	r.GET("/auth/login", authLoginHTML)
	r.POST("/auth/login", authLogin)

	r.GET("/", dashboardHTML)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
