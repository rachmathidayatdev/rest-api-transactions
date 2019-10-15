package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/rest-api-transaction/config"
	"github.com/rest-api-transaction/library"
	member "github.com/rest-api-transaction/members"
	repositoryTest "github.com/rest-api-transaction/members/repository/test"
	usecaseTest "github.com/rest-api-transaction/members/usecase/test"
	"github.com/rest-api-transaction/models"
)

//MemberHandler struct
type MemberHandler struct {
	MUsecase member.TestUsecase
}

//InitTesting func
func InitTesting() *MemberHandler {
	os.Setenv("API_PORT", "8080")
	os.Setenv("DB_HOST_LOCAL", "127.0.0.1")
	os.Setenv("DB_USERNAME_LOCAL", "root")
	os.Setenv("DB_PASS_LOCAL", "root")
	os.Setenv("DB_NAME_LOCAL", "learning")

	connection := config.GetConnection()
	MRepository := repositoryTest.TestNewMRepository(connection)
	MUsecase := usecaseTest.TestNewMUsecase(MRepository)

	handler := &MemberHandler{
		MUsecase: MUsecase,
	}

	return handler
}

//TestGetMember func
func TestGetMember(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/members/list?page=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]interface{}

	handlers := InitTesting()

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetMember)

	handler.ServeHTTP(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		_ = json.NewDecoder(res.Body).Decode(&data)
		t.Errorf("handler returned error: got %v want %v",
			data["got"], data["want"])
	}

	// Check the response body is what we expect.
	expected := map[string]interface{}{
		"message": "success",
		"data":    []models.Members{},
	}
	expectedString := library.MapStringToString(expected)

	if res.Body.String() != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expectedString)
	}
}

//GetMember func
func (m *MemberHandler) GetMember(w http.ResponseWriter, r *http.Request) {
	var offset int
	var limit int
	var totalData int

	limitTmp := r.FormValue("limit")
	pageTmp := r.FormValue("page")

	checkQueryParam := []string{"pageTmp", "limitTmp"}

	for _, item := range checkQueryParam {
		if item == "" {
			library.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"got":  fmt.Sprintf("%s null", item),
				"want": fmt.Sprintf("%s not null", item),
			})
			return
		}
	}

	limit, _ = strconv.Atoi(limitTmp)
	page, _ := strconv.Atoi(pageTmp)

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	payload := map[string]interface{}{
		"page":      page,
		"limit":     limit,
		"offset":    offset,
		"totalData": 0,
	}

	listMembers, totalData, error := m.MUsecase.TestGetMember(payload)

	if error != nil {
		var errostrings []string
		errostrings = append(errostrings, error.Error())
		library.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"got":  fmt.Sprintf("%s", "failed get data"),
			"want": fmt.Sprintf("%s", "success get data"),
		})
		return
	}

	fmt.Println(totalData)

	library.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    listMembers,
	})
}

//TestCreateMember func
func TestCreateMember(t *testing.T) {
	// req, err := http.NewRequest("POST", "/v1/members/create", bytes.NewBuffer([]byte(`{}`)))
	req, err := http.NewRequest("POST", "/v1/members/create",
		bytes.NewBuffer([]byte(`{
			"code": "asdasxcxzcsdf2",
			"name": "coba2",
			"email": "coba2@gmail.com",
			"password": "adasdaasdasdasd2",
			"phone": "01231231232",
			"gender": "laki laki",
			"religion": "islam"
		}`)))
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]interface{}

	handlers := InitTesting()

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateMember)

	handler.ServeHTTP(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		_ = json.NewDecoder(res.Body).Decode(&data)
		t.Errorf("handler returned error: got %v want %v",
			data["got"], data["want"])
	}
}

func (m *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	bodyData := make(map[string]interface{})

	decoder := json.NewDecoder(r.Body)

	if error := decoder.Decode(&bodyData); error != nil {
		library.ResponseJSON(w, http.StatusBadRequest, map[string]interface{}{
			"got":  "body data null",
			"want": "body data not null",
		})
		return
	}

	checkBodyParam := []string{"code", "name", "email", "password", "phone", "gender", "religion"}

	for _, item := range checkBodyParam {
		if bodyData[item] == nil {
			library.ResponseJSON(w, http.StatusBadRequest, map[string]interface{}{
				"got":  fmt.Sprintf("%s null", item),
				"want": fmt.Sprintf("%s not null", item),
			})
			return
		}
	}

	library.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
