package wallet

import (
	"database/sql"
)

type WalletRepository struct {
	Db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{
		Db: db,
	}
}

func (w *WalletRepository) Create(id string, request *NewWalletRequest) error {
	query := "INSERT INTO wallet(id, bank_name, description) VALUES($1, $2, $3)"
	_, err := w.Db.Exec(query, id, request.BankName, request.Description)

	return err
}

func (w *WalletRepository) DeleteById(id string) error {
	query := "DELETE FROM wallet WHERE id=$1"
	_, err := w.Db.Exec(query, id)

	return err
}
