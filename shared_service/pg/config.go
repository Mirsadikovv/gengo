package pg

import "gorm.io/gorm"

type ConnectionConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Port     int    `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSL_MODE"`
	TimeZone string `env:"DB_TIME_ZONE"`
}

type GormConfig = gorm.Config

var defaultConfig = ConnectionConfig{
	Host:     "localhost",
	User:     "postgres",
	Password: "postgres",
	DBName:   "postgres",
	Port:     5432,
	SSLMode:  "disable",
	TimeZone: "UTC",
}
