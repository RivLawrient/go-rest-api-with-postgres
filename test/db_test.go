package test

import (
	"fmt"
	"go-rest-api-with-postgres/internal/app/wallet"
	"go-rest-api-with-postgres/internal/config"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	vp := config.NewViper()

	db := config.GetConnection(vp)
	defer db.Close()

	err := db.Ping()

	assert.Nil(t, err)
}

func TestXxx(t *testing.T) {
	data := new(wallet.NewWalletRequest)

	field := reflect.ValueOf(*data)

	for i := 0; i < field.NumField(); i++ {

		fmt.Println(field.Type().Field(i).Tag.Get("json"))
	}

}
