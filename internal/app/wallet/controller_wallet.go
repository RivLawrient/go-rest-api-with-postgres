package wallet

import (
	"database/sql"
	"encoding/json"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
	"strings"
)

type WalletController struct {
	WalletUsecase *WalletUsecase
}

func NewWalletController(walletUsecase *WalletUsecase) *WalletController {
	return &WalletController{
		WalletUsecase: walletUsecase,
	}
}

func (w *WalletController) HandleNewWallet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	data := new(NewWalletRequest)

	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&data)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "body request is required",
		})
		return
	}

	if strings.TrimSpace(data.BankName) == "" {
		res.WriteHeader(http.StatusBadRequest)

		if data.BankName == "" {
			json.NewEncoder(res).Encode(model.WebResponse[string]{
				Errors: "field bank_name is required",
			})
			return
		}

		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "field bank_name cannot empty",
		})
		return
	}

	result, err := w.WalletUsecase.NewWallet(data)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(model.WebResponse[string]{
			Errors: "something error when insert data",
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(model.WebResponse[NewWalletResponse]{
		Data: *result,
	})
}

func (w *WalletController) HandleRemoveWallet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	//membaca params setelah "/wallet/"
	id := strings.Split(req.URL.Path, "/")[2]

	err := w.WalletUsecase.RemoveWallet(id)
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
		Data: "success remove wallet",
	})
}
