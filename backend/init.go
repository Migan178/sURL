package main

import (
	"github.com/Migan178/surl/routes"
	"github.com/Migan178/surl/routes/api"
	"github.com/Migan178/surl/routes/api/links"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()

	r.LoadHTMLGlob("static/templates/*.html")
	r.Static("/styles", "./static/styles")

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
