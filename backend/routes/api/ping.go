package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Migan178/surl/repository"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	preDBPing := time.Now()
	if err := repository.GetDatabase().Ping(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database doesn't work."})
		return
	}
	dbPing := time.Since(preDBPing).Milliseconds()

	c.JSON(http.StatusOK, gin.H{"db": dbPing})
}
