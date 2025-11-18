package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	var data repository.URL

	urn := c.Param("urn")

	row := repository.GetDatabase().QueryRow("select id, urn, redirect_url, created_at from urls where urn = ?;", urn)
	if err := row.Scan(&data.ID, &data.URN, &data.RedirectURL, &data.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}

		fmt.Println(err)
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
