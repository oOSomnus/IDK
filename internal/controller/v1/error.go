package v1

import (
	"I_Dev_Kit/internal/controller/v1/response"
	"encoding/json"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := response.Error{Error: msg}
	json.NewEncoder(w).Encode(resp)
}
