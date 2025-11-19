package links

import (
	"fmt"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	var data repository.CreateBody

	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to read post body"})
		return

	}

	createdData, err := repository.GetDatabase().CreateLink(data.RedirectURL)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured while inserting data"})
		return
	}

	c.JSON(http.StatusCreated, createdData)
}
