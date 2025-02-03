package transaction

import (
	"database/sql"
	"encoding/json"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
	"strings"
)

type TransactionController struct {
	TransactionUsecase *TransactionUsecase
}

func NewTransactionController(transactionUsecase *TransactionUsecase) *TransactionController {
	return &TransactionController{
		TransactionUsecase: transactionUsecase,
	}
}

func (t *TransactionController) HandleDetail(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	id := strings.Split(req.URL.Path, "/")[3]
	result, err := t.TransactionUsecase.Detail(id)
	if err != nil {
		if err == sql.ErrNoRows {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "id is not found",
			})
			return
		}

		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "something error when getting data",
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(model.WebResponse[TransactionDetailResponse]{
		Data: *result,
	})
}
