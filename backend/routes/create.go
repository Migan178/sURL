package routes

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"git.miganbox.com/migan/surl/repository"
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	var data repository.CreateBody

	if err := c.ShouldBind(&data); err != nil {
		code := http.StatusBadRequest
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	data.Normalize()
	if err := data.Validate(); err != nil {
		code := http.StatusBadRequest
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	createdData, err := repository.GetDatabase().CreateLink(c.Request.Context(), data.RedirectURL)
	if err != nil {
		if errors.Is(err, repository.ErrInvalid) {
			code := http.StatusBadRequest
			c.HTML(code, "error.html", gin.H{
				"status": http.StatusText(code),
				"code":   code,
			})
			return
		}

		slog.Error("failed to create short url", "err", err)
		code := http.StatusInternalServerError
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
