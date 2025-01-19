package test

import (
	"go-rest-api-with-postgres/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	db := config.GetConnection()
	defer db.Close()

	err := db.Ping()

	assert.Nil(t, err)
}
