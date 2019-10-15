package test

import (
	member "github.com/rest-api-transaction/members"
	"github.com/rest-api-transaction/models"
)

//TestMemberUsecase struct
type TestMemberUsecase struct {
	MRepository member.TestRepository
}

//TestNewMUsecase func
func TestNewMUsecase(m member.TestRepository) member.TestUsecase {
	return &TestMemberUsecase{
		MRepository: m,
	}
}

//TestGetMember func
func (m *TestMemberUsecase) TestGetMember(payload map[string]interface{}) ([]models.Members, int, error) {
	listMembers, totalData, error := m.MRepository.TestGetMember(payload)

	return listMembers, totalData, error
}
