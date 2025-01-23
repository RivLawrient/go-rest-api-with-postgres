package wallet

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type WalletRepository struct {
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{}
}

func (w *WalletRepository) Create(db *sql.DB, request *NewWalletRequest) {
	query := "INSERT INTO wallet(id, bank_name, description) VALUES($1, $2, $3)"
	result, err := db.Exec(query, uuid.New().String(), request.BankName, request.Description)

	if err != nil {
		log.Println(err)
	}

	log.Println(result.RowsAffected())
}
