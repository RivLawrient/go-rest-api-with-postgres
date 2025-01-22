package main

import (
	"go-rest-api-with-postgres/internal/config"
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
