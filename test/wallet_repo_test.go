package test

import (
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/config"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	req := wallet.NewWalletRequest{
		BankName:    "bca",
		Description: "tabungan",
	}

	err := wallet.NewWalletRepository(db).Create(uuid.New().String(), &req)

	assert.Nil(t, err)
}
