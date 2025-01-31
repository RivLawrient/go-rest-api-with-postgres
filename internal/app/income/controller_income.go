package income

import (
	"database/sql"
	"encoding/json"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
)

type IncomeController struct {
	IncomeUsecase *IncomeUsecase
}

func NewIncomeController(incomeUsecase *IncomeUsecase) *IncomeController {
	return &IncomeController{
		IncomeUsecase: incomeUsecase,
	}
}

func (i *IncomeController) HandleNewIncome(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	data := &NewIncomeRequest{}
	decode := json.NewDecoder(req.Body)

	if err := decode.Decode(&data); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "body request is required",
		})
		return
	}
	if data.WalletId == nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "wallet_id field is required",
		})
		return
	}

	result, err := i.IncomeUsecase.NewIncome(data)
	if err != nil {
		if err == sql.ErrNoRows {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "id is not found",
			})
			return
		}

		if err.Error() == "minus amount" {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "amount cannot be less than 0",
			})
			return
		}

		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "something error when create data",
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(model.WebResponse[NewIncomeResponse]{
		Data: *result,
	})
}
