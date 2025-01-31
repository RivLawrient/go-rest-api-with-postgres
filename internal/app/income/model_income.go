package income

import "time"

type NewIncomeRequest struct {
	Source   string  `json:"source"`
	Amount   int64   `json:"amount"`
	WalletId *string `json:"wallet_id"`
}

type NewIncomeResponse struct {
	Id        string    `json:"id"`
	Source    string    `json:"source"`
	Amount    int64     `json:"amount"`
	WalletId  string    `json:"wallet_id"`
	CreatedAt time.Time `json:"created_at"`
}
