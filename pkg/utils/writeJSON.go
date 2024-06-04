package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, value any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(value)
}
