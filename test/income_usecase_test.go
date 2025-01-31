package test

import (
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/config"
	"log"
	"testing"
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

	result, err := income.NewIncomeUsecase(income.NewIncomeRepository(db), wallet.NewWalletRepository(db)).NewIncome(&req)
	log.Println(result, err)
}
