package test

import (
	"go-rest-api-with-postgres/internal/app/income"
	"go-rest-api-with-postgres/internal/config"
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
