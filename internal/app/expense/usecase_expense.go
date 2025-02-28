package expense

import (
	"errors"
	"go-rest-api-with-postgres/internal/app/wallet"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ExpenseUsecase struct {
	ExpenseRepository *ExpenseRepository
	WalletRepository  *wallet.WalletRepository
}

func NewExpenseUsecase(expenseRepository *ExpenseRepository, walletRepository *wallet.WalletRepository) *ExpenseUsecase {
	return &ExpenseUsecase{
		ExpenseRepository: expenseRepository,
		WalletRepository:  walletRepository,
	}
}

func (e *ExpenseUsecase) New(request *NewExpenseRequest) (*NewExpenseResponse, error) {
	wallet, err := e.WalletRepository.FindById(*request.WalletId)
	if err != nil {
		return nil, err
	}

	if request.Quantity <= 0 {
		return nil, errors.New("0 quantity")
	}

	if request.Price < 0 {
		return nil, errors.New("minus price")
	}

	totalPrice := request.Price * int64(request.Quantity)
	if (wallet.Balance - totalPrice) <= 0 {
		return nil, errors.New("minus balance")
	}

	id := uuid.New().String()
	request.Item = strings.TrimSpace(request.Item)
	if err := e.ExpenseRepository.Add(id, request); err != nil {
		return nil, err
	}

	if err := e.WalletRepository.DecrementBalance(*request.WalletId, totalPrice); err != nil {
		return nil, err
	}
	return &NewExpenseResponse{
		Id:         id,
		Item:       request.Item,
		Quantity:   request.Quantity,
		Price:      request.Price,
		TotalPrice: request.Price * int64(request.Quantity),
		WalletId:   *request.WalletId,
		CreatedAt:  time.Now(),
	}, nil
}

