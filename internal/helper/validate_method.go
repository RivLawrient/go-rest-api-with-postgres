package helper

import (
	"encoding/json"
	"fmt"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
)

func ValidateMethod(res http.ResponseWriter, req *http.Request, method string) bool {
	if req.Method != method {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: fmt.Sprintf("method %s is not found", req.Method),
		})
		return true
	}
	return false
}
