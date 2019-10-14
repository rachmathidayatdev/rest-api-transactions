package config

import "os"

//Config struct
type Config struct {
	DB *DBConfig
}

//DBConfig struct
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Database string
	Host     string
	Port     string
	Charset  string
}

//GetConfigDB function
func GetConfigDB() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: os.Getenv("DB_USERNAME_LOCAL"),
			Password: os.Getenv("DB_PASS_LOCAL"),
			Database: os.Getenv("DB_NAME_LOCAL"),
			Host:     os.Getenv("DB_HOST_LOCAL"),
			Port:     "5432",
			Charset:  "utf8",
		},
	}
}
