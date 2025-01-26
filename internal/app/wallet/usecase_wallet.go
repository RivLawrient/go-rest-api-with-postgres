package wallet

import (
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
	err := w.WalletRepository.DeleteById(id)

	return err
}

func (w *WalletUsecase) ShowById(id string) (*ShowWalletResponse, error) {
	result, err := w.WalletRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result)

	return &ShowWalletResponse{
		Id:          id,
		BankName:    result.BankName,
		Description: result.Description,
		Balance:     result.Balance,
	}, nil
}

// func (w *WalletUsecase) ShowAll() (*[]Wallet, error) {

// }
