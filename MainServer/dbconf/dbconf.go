package dbconf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:mySQLLol#2v!@tcp(127.0.0.1:3306)/devjobs")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db
}

func DbConnToScrappy() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:fasd1423f@tcp(127.0.0.1:3306)/jobs")
	if err != nil {
		panic(err.Error())
	}
	return db
}
