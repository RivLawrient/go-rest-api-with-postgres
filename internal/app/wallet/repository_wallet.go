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

func (w *WalletRepository) FindById(id string) (*Wallet, error) {
	query := "SELECT bank_name, description, balance FROM wallet WHERE id=$1"
	result := w.Db.QueryRow(query, id)

	if result.Err() != nil {
		return nil, result.Err()
	}

	data := new(Wallet)
	if err := result.Scan(&data.BankName, &data.Description, &data.Balance); err != nil {
		return nil, err
	}

	return data, nil
}

func (w *WalletRepository) DeleteById(id string) error {
	query := "DELETE FROM wallet WHERE id=$1"
	result, err := w.Db.Exec(query, id)
	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return sql.ErrNoRows
	}

	return err
}
