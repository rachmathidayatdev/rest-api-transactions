package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//GetConnection function
func GetConnection() *gorm.DB {

	configDB := GetConfigDB()

	dbConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		configDB.DB.Username,
		configDB.DB.Password,
		configDB.DB.Host,
		configDB.DB.Port,
		configDB.DB.Database,
	)

	db, err := gorm.Open(configDB.DB.Dialect, dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
