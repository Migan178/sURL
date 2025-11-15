package links

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {
	var data repository.URL

	urn := c.Param("urn")

	row := repository.GetDatabase().QueryRow("select id, urn, redirect_url, created_at from urls where urn = ?;", urn)
	if err := row.Scan(&data.ID, &data.URN, &data.RedirectURL, &data.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s is not found", urn)})
			return
		}

		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get url"})
		return
	}

	c.JSON(http.StatusOK, data)
}
