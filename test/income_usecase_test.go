package test

import (
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/config"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIncome(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	wi := "e8624add-5577-4c29-b47d-6c5dcf4bdcef"
	req := income.NewIncomeRequest{
		Source:   "gaji",
		Amount:   -1,
		WalletId: &wi,
	}

	result, err := income.NewIncomeUsecase(income.NewIncomeRepository(db), wallet.NewWalletRepository(db)).New(&req)
	log.Println(result, err)
}

func Test(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	wi := "e8624add-5577-4c29-b47d-6c5dcf4bdcef"
	req := income.NewIncomeRequest{
		Source:   "gaji",
		Amount:   1,
		WalletId: &wi,
	}

	result, err := income.NewIncomeUsecase(income.NewIncomeRepository(db), wallet.NewWalletRepository(db)).New(&req)
	assert.Nil(t, err)

	err2 := income.NewIncomeUsecase(income.NewIncomeRepository(db), wallet.NewWalletRepository(db)).DeleteById(result.Id)
	assert.Nil(t, err2)
}
