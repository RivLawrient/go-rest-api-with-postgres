package expense

import "time"

type NewExpenseRequest struct {
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    int64   `json:"price"`
	WalletId *string `json:"wallet_id"`
}

type NewExpenseResponse struct {
	Id         string    `json:"id"`
	Item       string    `json:"item"`
	Quantity   int       `json:"quantity"`
	Price      int64     `json:"price"`
	TotalPrice int64     `json:"total_price"`
	WalletId   string    `json:"wallet_id"`
	CreatedAt  time.Time `json:"created_at"`
}
