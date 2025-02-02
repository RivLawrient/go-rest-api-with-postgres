package income

import (
	"database/sql"
	"encoding/json"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
	"strings"
)

type IncomeController struct {
	IncomeUsecase *IncomeUsecase
}

func NewIncomeController(incomeUsecase *IncomeUsecase) *IncomeController {
	return &IncomeController{
		IncomeUsecase: incomeUsecase,
	}
}

func (i *IncomeController) HandleNew(res http.ResponseWriter, req *http.Request) {
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

	result, err := i.IncomeUsecase.New(data)
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

func (i *IncomeController) HandleShowById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	//membaca params setelah "/wallet/"
	id := strings.Split(req.URL.Path, "/")[2]

	result, err := i.IncomeUsecase.ShowById(id)
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
	json.NewEncoder(res).Encode(model.WebResponse[ShowIncomeResponse]{
		Data: *result,
	})
}

func (i *IncomeController) HandleShowAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	result, err := i.IncomeUsecase.ShowALl()
	if err != nil {
		if err == sql.ErrNoRows {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "data is not found",
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
	json.NewEncoder(res).Encode(model.WebResponse[[]ShowIncomeResponse]{
		Data: *result,
	})
}

func (i *IncomeController) HandleDeleteById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	//membaca params setelah "/wallet/"
	id := strings.Split(req.URL.Path, "/")[2]

	err := i.IncomeUsecase.DeleteById(id)
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
			Errors: "something error when remove data",
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(model.WebResponse[string]{
		Data: "success remove income",
	})
}
