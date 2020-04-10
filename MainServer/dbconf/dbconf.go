package dbconf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:fasd1423f@tcp(127.0.0.1:3306)/devjobs")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db
}
