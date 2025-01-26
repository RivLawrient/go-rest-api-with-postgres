package test

import (
	"fmt"
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

func TestFindByIdSuccess(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	req := wallet.NewWalletRequest{
		BankName:    "bca",
		Description: "tabungan",
	}
	id := uuid.New().String()

	err := wallet.NewWalletRepository(db).Create(id, &req)
	assert.Nil(t, err)

	result, err := wallet.NewWalletRepository(db).FindById(id)
	assert.Nil(t, err)
	assert.Equal(t, result.BankName, req.BankName)
	assert.Equal(t, result.Description, req.Description)

	fmt.Println(id, result.BankName, result.Description, result.Balance)
}

func TestFindByIdFail(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)

	result, err := wallet.NewWalletRepository(db).FindById("random")
	assert.NotNil(t, err)
	assert.Nil(t, result)

	fmt.Println(err)
}

func TestDeleteByIdSuccess(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	req := wallet.NewWalletRequest{
		BankName:    "bca",
		Description: "tabungan",
	}
	id := uuid.New().String()

	err := wallet.NewWalletRepository(db).Create(id, &req)
	assert.Nil(t, err)

	err = wallet.NewWalletRepository(db).DeleteById(id)
	assert.Nil(t, err)
}

func TestDeleteByIdFail(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)

	err := wallet.NewWalletRepository(db).DeleteById("notfound")
	fmt.Println(err)
	assert.NotNil(t, err)
}
