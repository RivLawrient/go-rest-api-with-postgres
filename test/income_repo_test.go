package test

import (
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/config"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	wallet := "f0c58824-58f2-4c04-8a5e-f30ac3e69d4d"

	req := income.NewIncomeRequest{
		Source:   "gaji",
		Amount:   0,
		WalletId: &wallet,
	}

	err := income.NewIncomeRepository(db).Add(uuid.New().String(), &req)

	assert.Nil(t, err)
}
func TestFindById(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	id := "0d9e9bce-62da-485d-a817-f9a2091c4f0c"
	result, err := income.NewIncomeRepository(db).FindById(id)
	log.Println(result)
	assert.Nil(t, err)
}
func TestRemove(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	id := "a263f066-3d87-4f40-9312-2574bc68026a"
	err := income.NewIncomeRepository(db).RemoveById(id)

	assert.Nil(t, err)
}
func TestFindAllIncome(t *testing.T) {
	vp := TesViper()
	db := config.GetConnection(vp)
	result, err := income.NewIncomeRepository(db).FindAll()
	log.Println(result)
	assert.Nil(t, err)
}
