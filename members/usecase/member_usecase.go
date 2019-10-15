package usecase

import (
	member "github.com/rest-api-transaction/members"
	"github.com/rest-api-transaction/models"
)

//MemberUsecase struct
type MemberUsecase struct {
	MRepository member.Repository
}

//NewMUsecase func
func NewMUsecase(m member.Repository) member.Usecase {
	return &MemberUsecase{
		MRepository: m,
	}
}

//GetMember func
func (m *MemberUsecase) GetMember(payload map[string]interface{}) ([]models.Members, int, error) {
	listMembers, totalData, error := m.MRepository.GetMember(payload)

	return listMembers, totalData, error
}

//CreateMember func
func (m *MemberUsecase) CreateMember(payload map[string]interface{}) error {
	error := m.MRepository.CreateMember(payload)
	return error
}

//TestGetMember func
// func (m *MemberUsecase) TestGetMember(payload map[string]interface{}) ([]models.Members, int, error) {
// 	listMembers, totalData, error := m.MRepository.TestGetMember(payload)

// 	return listMembers, totalData, error
// }
