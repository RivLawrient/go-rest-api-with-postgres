package wallet

import "github.com/google/uuid"

type WalletUsecase struct {
	WalletRepository *WalletRepository
}

func NewWalletUsecase(walletRepository *WalletRepository) *WalletUsecase {
	return &WalletUsecase{
		WalletRepository: walletRepository,
	}
}

func (w *WalletUsecase) NewWallet(request *NewWalletRequest) error {
	id := uuid.New().String()

	err := w.WalletRepository.Create(id, request)

	return err
}
