package router

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rest-api-transaction/members/delivery/handler"
	"github.com/rest-api-transaction/members/repository"
	"github.com/rest-api-transaction/members/usecase"
	"github.com/rest-api-transaction/models"
)

//Route function
func Route(r *mux.Router, db *gorm.DB) {
	MRepository := repository.NewMRepository(db)
	MUsecase := usecase.NewMUsecase(MRepository)

	go models.DBMigrate(db)

	handler.InitMember(r, MUsecase)
}
