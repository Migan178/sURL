package links

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/Migan178/surl/utils"
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to read post body"})
		return
	}

	var data repository.CreateBody
	if err = json.Unmarshal(buf, &data); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong post body"})
		return
	}

	database := repository.GetDatabase()
	urn := utils.GetRandomString(20)

	row := database.QueryRow("select * from urls where urn = ?;", urn)
	if err = row.Err(); err != nil {
		if err == sql.ErrNoRows {
			urn = utils.GetRandomString(20)
		} else {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error while find duplicated urn"})
			return
		}
	}

	tx, err := database.Begin()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't being transaction"})
		return
	}

	_, err = tx.Exec("insert into urls(urn, redirect_url) values(?, ?);", urn, data.RedirectURL)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while insert data"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"urn": urn, "redirect_url": data.RedirectURL})
}
