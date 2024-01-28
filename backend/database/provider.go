package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DBProvider struct{}

func (p *DBProvider) OpenConnection() (*sql.DB, error) {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPWD"),
		Net:    "tcp",
		Addr:   os.Getenv("DBADDR"),
		DBName: os.Getenv("DBNAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
