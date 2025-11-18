package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Migan178/surl/repository"
	"github.com/Migan178/surl/utils"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateLinkWithForm(c *gin.Context) {
	var data repository.CreateBody

	if err := c.ShouldBind(&data); err != nil {
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return

	}

	database := repository.GetDatabase()
	urn := utils.GetRandomString(20)

	row := database.QueryRow("select * from urls where urn = ?;", urn)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			urn = utils.GetRandomString(20)
		} else {
			code := http.StatusInternalServerError
			fmt.Println(err)
			c.HTML(code, "error.html", gin.H{
				"status": http.StatusText(code),
				"code":   code,
			})
			return
		}
	}

	tx, err := database.Begin()
	if err != nil {
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	resp, err := tx.Exec("insert into urls(urn, redirect_url) values(?, ?);", urn, data.RedirectURL)
	if err != nil {
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	tx.Commit()

	var createdData repository.URL

	id, _ := resp.LastInsertId()
	row = database.QueryRow("select * from urls where id = ?;", id)
	if err = row.Scan(&createdData.ID, &createdData.URN, &createdData.RedirectURL, &createdData.CreatedAt); err != nil {
		code := http.StatusInternalServerError
		fmt.Println(err)
		c.HTML(code, "error.html", gin.H{
			"status": http.StatusText(code),
			"code":   code,
		})
		return
	}

	fmt.Printf("%+v\n", createdData)

	c.HTML(http.StatusOK, "info.html", gin.H{
		"urn":          createdData.URN,
		"redirect_url": createdData.RedirectURL,
		"created_at":   createdData.CreatedAt.Format("2006/01/02 15:04"),
	})
}
