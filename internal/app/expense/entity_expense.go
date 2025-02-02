package expense

import "time"

type Expense struct {
	Id         string
	Item       string
	Quantity   int
	Price      int64
	TotalPrice int64
	WalletId   string
	CreatedAt  time.Time
}
