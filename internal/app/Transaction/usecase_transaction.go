package transaction

import (
	"go-rest-api-with-postgres/internal/app/expense"
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/app/wallet"
)

type TransactionUsecase struct {
	WalletRepository  *wallet.WalletRepository
	IncomeRepository  *income.IncomeRepository
	ExpenseRepository *expense.ExpenseRepository
}

func NewTransactionUsecase(
	walletRepo *wallet.WalletRepository,
	incomeRepo *income.IncomeRepository,
	expenseRepo *expense.ExpenseRepository,
) *TransactionUsecase {
	return &TransactionUsecase{
		WalletRepository:  walletRepo,
		IncomeRepository:  incomeRepo,
		ExpenseRepository: expenseRepo,
	}
}

func (t *TransactionUsecase) Detail(idWallet string) (*TransactionDetailResponse, error) {
	walletResult, err := t.WalletRepository.FindById(idWallet)
	if err != nil {
		return nil, err
	}

	incomeResult, err := t.IncomeRepository.FindAllByWalleId(idWallet)
	if err != nil {
		return nil, err
	}

	var totalAmountIncome int64 = 0
	incResponses := []income.ShowIncomeResponse{}
	for _, data := range *incomeResult {
		response := &income.ShowIncomeResponse{
			Id:        data.Id,
			Source:    data.Source,
			Amount:    data.Amount,
			WalletId:  data.WalletId,
			CreatedAt: data.CreatedAt,
		}

		totalAmountIncome = totalAmountIncome + data.Amount

		incResponses = append(incResponses, *response)
	}

	expenseResult, err := t.ExpenseRepository.FindAllByWalletId(idWallet)
	if err != nil {
		return nil, err
	}

	totalQuantityExpense := 0
	var totalAmountExpense int64 = 0
	expResponses := []expense.ShowExpenseResponse{}
	for _, data := range *expenseResult {
		response := &expense.ShowExpenseResponse{
			Id:         data.Id,
			Item:       data.Item,
			Quantity:   data.Quantity,
			Price:      data.Price,
			TotalPrice: data.TotalPrice,
			WalletId:   data.WalletId,
			CreatedAt:  data.CreatedAt,
		}
		totalQuantityExpense = totalQuantityExpense + data.Quantity
		totalAmountExpense = totalAmountExpense + data.TotalPrice

		expResponses = append(expResponses, *response)
	}

	totalItemExpense := len(expResponses)

	return &TransactionDetailResponse{
		Wallet: (*wallet.ShowWalletResponse)(walletResult),
		Overall: &OverallResponse{
			TotalIncome:          totalAmountIncome,
			TotalItemExpense:     totalItemExpense,
			TotalQuantityExpense: totalQuantityExpense,
			TotalExpense:         totalAmountExpense,
		},
		Income:  &incResponses,
		Expense: &expResponses,
	}, err

}
