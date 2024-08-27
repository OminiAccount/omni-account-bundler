package services

import (
	"net/http"
)

func corsAndAllowMethod(w *http.ResponseWriter, request *http.Request) bool {
	if request.Method == http.MethodOptions {
		return false
	}
	if request.Method != http.MethodPost {
		http.Error(*w, "Method not allowed.", http.StatusMethodNotAllowed)
		return false
	}
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	(*w).Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Content-type", "application/json")
	return true
}
