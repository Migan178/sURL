package main

import (
	"github.com/Migan178/surl/routes"
	"github.com/Migan178/surl/routes/links"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()

	r.GET("/", routes.Root)
	r.GET("/ping", routes.Ping)

	linkRouter := r.Group("/links")
	{
		linkRouter.POST("/", links.CreateLink)
	}
}
