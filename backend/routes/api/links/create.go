package links

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/Migan178/surl/utils"
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	var data repository.CreateBody

	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to read post body"})
		return

	}

	database := repository.GetDatabase()
	urn := utils.GetRandomString(20)

	row := database.QueryRow("select * from urls where urn = ?;", urn)
	if err := row.Err(); err != nil {
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

	resp, err := tx.Exec("insert into urls(urn, redirect_url) values(?, ?);", urn, data.RedirectURL)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while insert data"})
		return
	}

	tx.Commit()

	var createdData repository.URL

	id, _ := resp.LastInsertId()
	row = database.QueryRow("select * from urls where id = ?;", id)
	if err = row.Scan(&createdData.ID, &createdData.URN, &createdData.RedirectURL, &createdData.CreatedAt); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while insert data"})
		return
	}

	fmt.Printf("%+v\n", createdData)

	c.JSON(http.StatusCreated, createdData)
}
