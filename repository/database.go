package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"git.miganbox.com/migan/surl/backend/utils"
	"git.miganbox.com/migan/surl/configs"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type SURLDatabase struct {
	*sql.DB
}

var databaseInstance *SURLDatabase

func GetDatabase() *SURLDatabase {
	if databaseInstance == nil {
		config := configs.GetConfig().Database
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
	var createdData URL

	for {
		urn := utils.GetRandomString(20)
		createdRow := d.QueryRow("insert into urls(urn, redirect_url) values(?, ?) returning id, urn, redirect_url, created_at", urn, redirectURL)

		if err := createdRow.Scan(&createdData.ID, &createdData.URN, &createdData.RedirectURL, &createdData.CreatedAt); err != nil {
			if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok {
				if mysqlErr.Number == 1062 {
					continue
				}

				return nil, err
			}

			return nil, err
		}

		return &createdData, nil
	}
}

func (d *SURLDatabase) Close() error {
	return d.DB.Close()
}
