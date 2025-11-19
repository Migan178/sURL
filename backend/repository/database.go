package repository

import (
	"database/sql"
	"fmt"

	"github.com/Migan178/surl/configs"
	"github.com/Migan178/surl/utils"
	_ "github.com/go-sql-driver/mysql"
)

type SURLDatabase struct {
	*sql.DB
}

var databaseInstance *SURLDatabase

func GetDatabase() *SURLDatabase {
	if databaseInstance == nil {
		config := configs.GetConfigs().Database
		conn, err := sql.Open(
			"mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?parseTime=true",
				config.Username,
				config.Password,
				config.Hostname,
				config.Port,
				config.DatabaseName,
			))
		if err != nil {
			panic(err)
		}

		databaseInstance = &SURLDatabase{conn}
	}

	return databaseInstance
}

func (d *SURLDatabase) CreateLink(redirectURL string) (*URL, error) {
	urn := utils.GetRandomString(20)

	row := d.QueryRow("select * from urls where urn = ?;", urn)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			urn = utils.GetRandomString(20)
		} else {
			return nil, err
		}
	}

	tx, err := d.Begin()
	if err != nil {
		return nil, err
	}

	resp, err := tx.Exec("insert into urls(urn, redirect_url) values(?, ?);", urn, redirectURL)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	var createdData URL

	id, _ := resp.LastInsertId()
	row = d.QueryRow("select * from urls where id = ?;", id)
	if err = row.Scan(&createdData.ID, &createdData.URN, &createdData.RedirectURL, &createdData.CreatedAt); err != nil {
		return nil, err
	}

	return &createdData, nil
}
