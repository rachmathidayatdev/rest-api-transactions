package member

import "github.com/rest-api-transaction/models"

//Usecase interface
type Usecase interface {
	GetMember(payload map[string]interface{}) ([]models.Members, int, error)
	CreateMember(payload map[string]interface{}) error
}

// TestUsecase func
type TestUsecase interface {
	TestGetMember(payload map[string]interface{}) ([]models.Members, int, error)
}
