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

func (w *WalletRepository) FindAll() (*[]Wallet, error) {
	query := "SELECT id, bank_name, description, balance FROM wallet"
	result, err := w.Db.Query(query)
	if err != nil {
		return nil, err
	}
	datas := []Wallet{}

	for result.Next() {
		data := Wallet{}
		err := result.Scan(&data.Id, &data.BankName, &data.Description, &data.Balance)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	if len(datas) == 0 {
		return nil, sql.ErrNoRows
	}
	return &datas, nil
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

func (w *WalletRepository) UpdateById(id string, request *EditWalletRequest) error {
	query := "UPDATE wallet SET bank_name = $1, description=$2 WHERE id=$3"
	_, err := w.Db.Exec(query, request.BankName, request.Description, id)

	return err
}

func (w *WalletRepository) IncrementBalance(id string, amount int64) error {
	query := "UPDATE wallet SET balance = balance + $1 WHERE id=$2"
	_, err := w.Db.Exec(query, amount, id)

	return err
}
