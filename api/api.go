package api

import (
	"encoding/json"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(struct {

	}{})
}