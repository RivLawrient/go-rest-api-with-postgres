package router

import (
	"encoding/json"
	"fmt"
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/model"
	"net/http"
)

// membungkus fungsi Handle di method yang dibutuhkan
type MethodHandlers struct {
	get    func(http.ResponseWriter, *http.Request)
	post   func(http.ResponseWriter, *http.Request)
	put    func(http.ResponseWriter, *http.Request)
	delete func(http.ResponseWriter, *http.Request)
}

// menggabungkan fungsi Handle jika berada di endpoint yang sama namun method yang berbeda
func (mh *MethodHandlers) Handle(w http.ResponseWriter, r *http.Request) {
	handler := map[string]func(http.ResponseWriter, *http.Request){
		http.MethodGet:    mh.get,
		http.MethodPost:   mh.post,
		http.MethodPut:    mh.put,
		http.MethodDelete: mh.delete,
	}[r.Method]

	if handler != nil {
		handler(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.WebResponse[string]{
			Errors: fmt.Sprintf("method %s is not found", r.Method),
		})
	}
}

type RouterConfig struct {
	Routing          *http.ServeMux
	WalletController *wallet.WalletController
	IncomeController *income.IncomeController
}

// membungkus semua endpoint yang dibuat.
// menghindari penggunaan root endpoint ("/") agar menghasilkan 404 page not found.
func (c *RouterConfig) Route() {

	walletHandle := &MethodHandlers{
		post: c.WalletController.HandleNewWallet,
		get:  c.WalletController.HandleShowAll,
	}
	c.Routing.HandleFunc("/wallet", walletHandle.Handle)

	walletParamHandle := &MethodHandlers{
		delete: c.WalletController.HandleRemoveWallet,
		get:    c.WalletController.HandleShowById,
		put:    c.WalletController.HandleEditWalletById,
	}
	c.Routing.HandleFunc("/wallet/", walletParamHandle.Handle)

	incomeHandle := &MethodHandlers{
		post: c.IncomeController.HandleNewIncome,
	}
	c.Routing.HandleFunc("/income", incomeHandle.Handle)

}
