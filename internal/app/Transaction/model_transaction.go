package transaction

import (
	"go-rest-api-with-postgres/internal/app/expense"
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/app/wallet"
)

type TransactionDetailResponse struct {
	Wallet  *wallet.ShowWalletResponse     `json:"wallet"`
	Overall *OverallResponse               `json:"overall"`
	Income  *[]income.ShowIncomeResponse   `json:"income"`
	Expense *[]expense.ShowExpenseResponse `json:"expense"`
}

type OverallResponse struct {
	TotalIncome    int64 `json:"total_income"`
	TotalItemExpense     int   `json:"total_item_expense"`
	TotalQuantityExpense int   `json:"total_quantity_expense"`
	TotalExpense   int64 `json:"total_expense"`
}
