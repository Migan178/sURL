package routes

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"git.miganbox.com/migan/surl/repository"
	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	urn := c.Param("urn")

	data, err := repository.GetDatabase().Find(urn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}

		code := http.StatusInternalServerError
		slog.Error("failed to redirect", "err", err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, data.RedirectURL)
}
