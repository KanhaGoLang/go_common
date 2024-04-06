package common

import (
	"database/sql"
	"log"

	"github.com/fatih/color"
)

var MyLogger = log.New(color.Output, "", log.Ldate|log.Ltime|log.Llongfile)

func NewDatabaseConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "sanjeev"
	dbAddr := "tcp(localhost:3306)"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbAddr+"/"+dbName)
	if err != nil {
		return nil, err
	}

	MyLogger.Println(color.GreenString("DB connected"))

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
