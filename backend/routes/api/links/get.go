package links

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"git.miganbox.com/migan/surl/repository"
	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {
	urn := c.Param("urn")

	data, err := repository.GetDatabase().Find(urn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s is not found", urn)})
			return
		}

		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get url"})
		return
	}

	c.JSON(http.StatusOK, data)
}
