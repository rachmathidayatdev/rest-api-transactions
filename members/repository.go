package member

import "github.com/rest-api-transaction/models"

//Repository interface
type Repository interface {
	GetMember(payload map[string]interface{}) ([]models.Members, int, error)
	CreateMember(payload map[string]interface{}) error
}
