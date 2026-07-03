package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	MySQLDriver = "mysql"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Database string `koanf:"database"`
}

func (c Config) GetDSN() string {
	DSN := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", c.Username, c.Password, c.Host, c.Port, c.Database)
	return DSN
}

type Connection struct {
	DB *sql.DB
}

func New(config Config) Connection {
	db, err := sql.Open(MySQLDriver, config.GetDSN())
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return Connection{
		DB: db,
	}
}
