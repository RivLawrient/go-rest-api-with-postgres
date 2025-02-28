package wallet

type NewWalletRequest struct {
	BankName    string `json:"bank_name"`
	Description string `json:"description"`
}

type NewWalletResponse struct {
	Id          string `json:"id"`
	BankName    string `json:"bank_name"`
	Description string `json:"description"`
	Balance     int64  `json:"balance"`
}

type ShowWalletResponse struct {
	Id          string `json:"id"`
	BankName    string `json:"bank_name"`
	Description string `json:"description"`
	Balance     int64  `json:"balance"`
}

type EditWalletRequest struct {
	BankName    *string `json:"bank_name"`
	Description string  `json:"description"`
}

type EditWalletResponse struct {
	Id          string `json:"id"`
	BankName    string `json:"bank_name"`
	Description string `json:"description"`
}
