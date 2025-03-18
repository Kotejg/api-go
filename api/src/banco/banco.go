package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

func Conectar() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", config.StringConexaoDB)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
