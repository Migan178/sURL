package routes

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"git.miganbox.com/migan/surl/repository"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	urn := c.Param("urn")

	data, err := repository.GetDatabase().Find(urn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}

		slog.Error("failed to get info", "err", err)
		code := http.StatusInternalServerError
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	c.HTML(http.StatusOK, "info.html", gin.H{
		"urn":          data.URN,
		"redirect_url": data.RedirectURL,
		"created_at":   data.CreatedAt.Format("2006/01/02 15:04"),
	})
}
