package routes

import (
	"fmt"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateLinkWithForm(c *gin.Context) {
	var data repository.CreateBody

	if err := c.ShouldBind(&data); err != nil {
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	createdData, err := repository.GetDatabase().CreateLink(data.RedirectURL)
	if err != nil {
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	c.HTML(http.StatusOK, "info.html", gin.H{
		"urn":          createdData.URN,
		"redirect_url": createdData.RedirectURL,
		"created_at":   createdData.CreatedAt.Format("2006/01/02 15:04"),
	})
}
