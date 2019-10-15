package test

import (
	"github.com/jinzhu/gorm"
	member "github.com/rest-api-transaction/members"
	"github.com/rest-api-transaction/models"
)

//DBHandler struct
type DBHandler struct {
	DB *gorm.DB
}

//TestNewMRepository func
func TestNewMRepository(DB *gorm.DB) member.TestRepository {
	return &DBHandler{DB}
}

//TestGetMember func
func (connection *DBHandler) TestGetMember(payload map[string]interface{}) ([]models.Members, int, error) {
	var error error
	var totalData int
	return []models.Members{}, totalData, error
}
