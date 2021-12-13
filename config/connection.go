package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	cfg, err := ConfigYAML()
	db, err := sql.Open(cfg.Database, cfg.Username+":"+cfg.Password+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.DBName)
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(cfg.MaxConnection)
	db.SetMaxIdleConns(cfg.MaxIdleConnection)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return db
}
