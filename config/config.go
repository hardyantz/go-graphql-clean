package config

import "database/sql"

type Config struct {
	DB *sql.DB
}

func InitConfig() Config {
	dbConn := ConnDB()
	return Config{DB: dbConn}
}
