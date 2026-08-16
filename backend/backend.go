package backend

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"git.miganbox.com/migan/surl/backend/routes"
	"git.miganbox.com/migan/surl/backend/routes/api"
	"git.miganbox.com/migan/surl/backend/routes/api/links"
	"git.miganbox.com/migan/surl/configs"
	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFS embed.FS

func New() *http.Server {
	r := gin.Default()

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

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", configs.GetConfig().Backend.Port),
		Handler: r,
	}
}
