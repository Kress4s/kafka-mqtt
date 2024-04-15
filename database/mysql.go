package database

import (
	"database/sql"
	"fmt"
	"kafkaMq/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func newMysqlDriver(cfg *config.Configure) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Database.Username, cfg.Database.Password,
		cfg.Database.Host, cfg.Database.Port, cfg.Database.DB)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
	}
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Database.MaxIdleConns)
	log.Println("database connect successful...")
	return db
}
