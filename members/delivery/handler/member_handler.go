package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rest-api-transaction/library"
	member "github.com/rest-api-transaction/members"
)

//MemberHandler struct
type MemberHandler struct {
	MUsecase member.Usecase
}

//InitMember func
func InitMember(r *mux.Router, MUsecase member.Usecase) {
	handler := &MemberHandler{
		MUsecase: MUsecase,
	}

	r.Handle("/v1/members/list",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.GetMember(w, r)
		}),
	).Methods("GET")

	r.Handle("/v1/members/create",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.CreateMember(w, r)
		}),
	).Methods("POST")
}

//GetMember func
func (m *MemberHandler) GetMember(w http.ResponseWriter, r *http.Request) {
	var offset int
	var limit int
	var totalData int

	limitTmp := r.FormValue("limit")
	limit, _ = strconv.Atoi(limitTmp)

	pageTmp := r.FormValue("page")
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

	listMembers, totalData, error := m.MUsecase.GetMember(payload)

	if error != nil {
		var errostrings []string
		errostrings = append(errostrings, error.Error())
		library.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "failed",
			"error":   errostrings,
		})
		return
	}

	fmt.Println(totalData)

	library.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    listMembers,
	})
}

//CreateMember func
func (m *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	bodyData := make(map[string]interface{})

	decoder := json.NewDecoder(r.Body)

	if error := decoder.Decode(&bodyData); error != nil {
		library.ResponseJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"error":   error.Error(),
		})
		return
	}

	error := m.MUsecase.CreateMember(bodyData)

	if error != nil {
		library.ResponseJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"error":   error.Error(),
		})
		return
	}

	library.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    bodyData,
	})
}
