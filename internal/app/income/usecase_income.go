package income

import (
	"errors"
	"go-rest-api-with-postgres/internal/app/wallet"
	"strings"
	"time"

	"github.com/google/uuid"
)

type IncomeUsecase struct {
	IncomeRepository *IncomeRepository
	WalletRepository *wallet.WalletRepository
}

func NewIncomeUsecase(incomeRepository *IncomeRepository, walletRepository *wallet.WalletRepository) *IncomeUsecase {
	return &IncomeUsecase{
		IncomeRepository: incomeRepository,
		WalletRepository: walletRepository,
	}
}

func (i *IncomeUsecase) New(request *NewIncomeRequest) (*NewIncomeResponse, error) {

	//check wallet_id
	if _, err := i.WalletRepository.FindById(*request.WalletId); err != nil {
		return nil, err
	}

	//amount tidak boleh kurang dari 0
	if request.Amount < 0 {
		return nil, errors.New("minus amount")
	}

	id := uuid.New().String()
	request.Source = strings.TrimSpace(request.Source)

	if err := i.IncomeRepository.Add(id, request); err != nil {
		return nil, err
	}

	if err := i.WalletRepository.IncrementBalance(*request.WalletId, request.Amount); err != nil {
		return nil, err
	}

	return &NewIncomeResponse{
		Id:        id,
		Source:    request.Source,
		Amount:    request.Amount,
		WalletId:  *request.WalletId,
		CreatedAt: time.Now(),
	}, nil
}
func (i *IncomeUsecase) ShowById(id string) (*ShowIncomeResponse, error) {
	result, err := i.IncomeRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return &ShowIncomeResponse{
		Id:        result.Id,
		Source:    result.Source,
		Amount:    result.Amount,
		WalletId:  result.WalletId,
		CreatedAt: result.CreatedAt,
	}, nil
}

func (i *IncomeUsecase) ShowALl() (*[]ShowIncomeResponse, error) {
	result, err := i.IncomeRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := []ShowIncomeResponse{}

	for _, data := range *result {
		response := &ShowIncomeResponse{
			Id:        data.Id,
			Source:    data.Source,
			Amount:    data.Amount,
			WalletId:  data.WalletId,
			CreatedAt: data.CreatedAt,
		}

		responses = append(responses, *response)
	}

	return &responses, nil
}

func (i *IncomeUsecase) DeleteById(id string) error {
	//check wallet_id
	if _, err := i.IncomeRepository.FindById(id); err != nil {
		return err
	}

	err := i.IncomeRepository.RemoveById(id)

	return err
}
