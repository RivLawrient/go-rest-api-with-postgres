package test

import (
	"go-rest-api-with-postgres/internal/app/expense"
	"go-rest-api-with-postgres/internal/config"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddExpense(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	wallet := "f0c58824-58f2-4c04-8a5e-f30ac3e69d4d"

	req := expense.NewExpenseRequest{
		Item:     "somethi",
		Quantity: 12,
		Price:    20,
		WalletId: &wallet,
	}

	err := expense.NewExpenseRepository(db).Add(uuid.New().String(), &req)
	assert.Nil(t, err)
}

func TestFindByIdExpense(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	id := "efa62118-dcba-4e9b-a2c9-e9e679b28092"
	result, err := expense.NewExpenseRepository(db).FindById(id)
	log.Println(result)
	assert.Nil(t, err)
}
func TestFindAllExpense(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	result, err := expense.NewExpenseRepository(db).FindAll()
	for _, data := range *result {
		log.Println(data)
	}
	assert.Nil(t, err)
}
