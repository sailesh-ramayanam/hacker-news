package dbhn

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectToDB() error {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "hacker_news",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connected to database")
	return nil
}
