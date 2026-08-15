package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/Migan178/surl/routes"
	"github.com/Migan178/surl/routes/api"
	"github.com/Migan178/surl/routes/api/links"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

//go:embed static/*
var staticFS embed.FS

func init() {
	r = gin.Default()

	tmpl, err := template.ParseFS(staticFS, "static/templates/*.html")
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(tmpl)

	stylesFS, err := fs.Sub(staticFS, "static/styles")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/styles", http.FS(stylesFS))

	scriptsFS, err := fs.Sub(staticFS, "static/scripts")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/scripts", http.FS(scriptsFS))

	r.GET("/", routes.Home)
	r.GET("/:urn", routes.Redirect)
	r.GET("/info/:urn", routes.Info)

	r.POST("/", routes.CreateLinkWithForm)

	apiRouter := r.Group("/api")
	{
		r.GET("/ping", api.Ping)

		linkRouter := apiRouter.Group("/links")
		{
			linkRouter.POST("/", links.CreateLink)
			linkRouter.GET("/:urn", links.GetLink)
		}
	}
}
