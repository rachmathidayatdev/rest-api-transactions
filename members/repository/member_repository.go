package repository

import (
	"github.com/jinzhu/gorm"
	member "github.com/rest-api-transaction/members"
	"github.com/rest-api-transaction/models"
)

//DBHandler struct
type DBHandler struct {
	DB *gorm.DB
}

//NewMRepository func
func NewMRepository(DB *gorm.DB) member.Repository {
	return &DBHandler{DB}
}

//GetMember func
func (connection *DBHandler) GetMember(payload map[string]interface{}) ([]models.Members, int, error) {
	var totalData int
	members := []models.Members{}

	error := connection.DB.Find(&members).Count(&totalData).Error

	error = connection.DB.Limit(payload["limit"].(int)).Offset(payload["offset"].(int)).Find(&members).Error

	return members, totalData, error
}

//CreateMember func
func (connection *DBHandler) CreateMember(payload map[string]interface{}) error {
	member := models.Members{
		Code:     payload["code"].(string),
		Name:     payload["name"].(string),
		Email:    payload["email"].(string),
		Password: payload["password"].(string),
		Phone:    payload["phone"].(string),
		Gender:   payload["gender"].(string),
		Religion: payload["religion"].(string),
	}

	tx := connection.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
