package web

import (
	"encoding/json"
	"github.com/IuryAlves/go-server/internal"
	"io/ioutil"
	"net/http"
	"strconv"
)

func TODOHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(internal.ListTODOs())
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		todo := new(internal.Todo)
		json.Unmarshal(reqBody, &todo)
		internal.SaveTODO(todo)
	case "DELETE":
		todoId, _ := strconv.Atoi(r.URL.Query().Get(":id"))
		internal.DeleteTODO(todoId)
	default:
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	}
}