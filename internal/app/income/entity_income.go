package income

import "time"

type Income struct {
	Id        string
	Source    string
	Amount    int64
	WalletId  string
	CreatedAt time.Time
}
