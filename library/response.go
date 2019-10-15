package library

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON makes the response with payload as json format
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

//MapStringToString func
func MapStringToString(payload map[string]interface{}) string {
	var expectedString string
	expectedData, err := json.Marshal(payload)

	if err != nil {
		return expectedString
	}

	expectedString = string(expectedData)

	return expectedString
}
