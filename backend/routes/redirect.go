package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"git.miganbox.com/migan/surl/repository"
	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	var data repository.URL

	urn := c.Param("urn")

	row := repository.GetDatabase().QueryRow("select id, urn, redirect_url, created_at from urls where urn = ?;", urn)
	if err := row.Scan(&data.ID, &data.URN, &data.RedirectURL, &data.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, data.RedirectURL)
}
