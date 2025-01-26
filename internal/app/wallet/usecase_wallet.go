package wallet

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
)

type WalletUsecase struct {
	WalletRepository *WalletRepository
}

func NewWalletUsecase(walletRepository *WalletRepository) *WalletUsecase {
	return &WalletUsecase{
		WalletRepository: walletRepository,
	}
}

func (w *WalletUsecase) NewWallet(request *NewWalletRequest) (*NewWalletResponse, error) {
	id := uuid.New().String()
	request.BankName = strings.TrimSpace(request.BankName)
	err := w.WalletRepository.Create(id, request)
	if err != nil {
		return nil, err
	}

	return &NewWalletResponse{
		Id:          id,
		BankName:    request.BankName,
		Description: request.Description,
		Balance:     0,
	}, nil
}

func (w *WalletUsecase) RemoveWallet(id string) error {
	_, err := w.WalletRepository.FindById(id)
	if err != sql.ErrNoRows {
		return err
	}
	return err
}
