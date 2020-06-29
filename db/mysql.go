package db

import (
	"app/config"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySql(config *config.DBConfig) (*sql.DB, error) {
	loc := strings.Replace(config.TimeZone, "/", "%2F", 1)
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", config.User, config.Password, config.Host, config.Port, config.DBName, loc)
	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
