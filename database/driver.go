package database

import (
	"database/sql"
	"kafkaMq/config"
	"sync"
)

var (
	once     sync.Once
	instance *sql.DB
)

func GetMysqlDBDriver() *sql.DB {
	once.Do(func() {
		cfg := config.GetConfigure()
		switch cfg.Database.Type {
		case "mysql":
			instance = newMysqlDriver(cfg)
		default:
			instance = nil
		}
	})
	return instance
}
