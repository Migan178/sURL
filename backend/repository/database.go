package repository

import (
	"database/sql"
	"fmt"

	"github.com/Migan178/surl/configs"
	_ "github.com/go-sql-driver/mysql"
)

var databaseInstance *sql.DB

func GetDatabase() *sql.DB {
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

		databaseInstance = conn
	}

	return databaseInstance
}
