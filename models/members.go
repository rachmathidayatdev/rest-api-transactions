package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Members struct
type Members struct {
	ID        int        `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Gender    string     `json:"gender"`
	Religion  string     `json:"religion"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	CreatedBy string     `json:"created_by,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy string     `json:"deleted_by,omitempty" gorm:"DEFAULT:null"`
}

//DBMigrate migrate
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&Members{},
	)

	return db
}
