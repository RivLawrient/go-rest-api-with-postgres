package main

import (
	"fmt"
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/config"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	vp := config.NewViper()

	db := config.GetConnection(vp)
	defer db.Close()

	err := db.Ping()

	assert.Nil(t, err)
}

func TestRepo(t *testing.T) {
	vp := config.NewViper()

	group := &sync.WaitGroup{}

	data := &wallet.NewWalletRequest{
		BankName:    "some",
		Description: "yes",
	}

	for i := 0; i < 101; i++ {
		group.Add(1)
		go func() {
			db := config.GetConnection(vp)
			wallet.NewWalletRepository().Create(db, uuid.New().String(), data)
		}()

	}

	group.Wait()
	fmt.Println("========SELESAI=======")
}
