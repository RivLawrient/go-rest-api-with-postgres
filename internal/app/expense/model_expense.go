package expense

type NewExpenseRequest struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
	Price    int64  `json:"price"`
	WalletId string `json:"wallet_id"`
}
