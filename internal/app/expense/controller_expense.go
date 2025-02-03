package expense

import (
	"database/sql"
	"encoding/json"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
)

type ExpenseController struct {
	ExpenseUsecase *ExpenseUsecase
}

func NewExpenseController(expenseUsecase *ExpenseUsecase) *ExpenseController {
	return &ExpenseController{
		ExpenseUsecase: expenseUsecase,
	}
}

func (e *ExpenseController) HandleNew(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	data := &NewExpenseRequest{}
	decode := json.NewDecoder(req.Body)

	if err := decode.Decode(data); err != nil {
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

	result, err := e.ExpenseUsecase.New(data)
	if err != nil {
		if err == sql.ErrNoRows {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "id is not found",
			})
			return
		}

		if err.Error() == "0 quantity" {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "quantity cannot 0 or less",
			})
			return
		}

		if err.Error() == "minus price" {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "price cannot be less than 0",
			})
			return
		}

		if err.Error() == "minus balance" {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "your wallet balance not enough",
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
	json.NewEncoder(res).Encode(model.WebResponse[NewExpenseResponse]{
		Data: *result,
	})
}
