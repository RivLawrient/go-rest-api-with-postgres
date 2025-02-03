package income

import (
	"database/sql"
)

type IncomeRepository struct {
	Db *sql.DB
}

func NewIncomeRepository(db *sql.DB) *IncomeRepository {
	return &IncomeRepository{
		Db: db,
	}
}

func (i *IncomeRepository) Add(id string, request *NewIncomeRequest) error {
	query := "INSERT INTO income(id, source, amount, wallet_id) VALUES($1, $2, $3, $4)"
	_, err := i.Db.Exec(query, id, request.Source, request.Amount, request.WalletId)

	return err
}

func (i *IncomeRepository) FindById(id string) (*Income, error) {
	query := "SELECT id, source, amount, wallet_id, created_at FROM income WHERE id=$1"
	result := i.Db.QueryRow(query, id)

	if result.Err() != nil {
		return nil, result.Err()
	}

	data := new(Income)
	if err := result.Scan(&data.Id, &data.Source, &data.Amount, &data.WalletId, &data.CreatedAt); err != nil {
		return nil, err
	}

	return data, nil
}

func (i *IncomeRepository) FindAll() (*[]Income, error) {
	query := "SELECT id, source, amount, wallet_id, created_at FROM income"
	result, err := i.Db.Query(query)
	if err != nil {
		return nil, err
	}
	list := []Income{}
	for result.Next() {
		data := Income{}
		err := result.Scan(&data.Id, &data.Source, &data.Amount, &data.WalletId, &data.CreatedAt)
		if err != nil {
			return nil, err
		}

		list = append(list, data)
	}

	if len(list) == 0 {
		return nil, sql.ErrNoRows
	}

	return &list, nil
}

func (i *IncomeRepository) RemoveById(id string) error {
	query := "DELETE FROM income WHERE id=$1"
	_, err := i.Db.Exec(query, id)

	return err
}

func (i *IncomeRepository) FindAllByWalleId(idWallet string) (*[]Income, error) {
	query := "SELECT id, source, amount, wallet_id, created_at FROM income WHERE wallet_id=$1"
	result, err := i.Db.Query(query, idWallet)
	if err != nil {
		return nil, err
	}
	list := []Income{}
	for result.Next() {
		data := Income{}
		err := result.Scan(&data.Id, &data.Source, &data.Amount, &data.WalletId, &data.CreatedAt)
		if err != nil {
			return nil, err
		}

		list = append(list, data)
	}

	if len(list) == 0 {
		return nil, sql.ErrNoRows
	}

	return &list, nil
}
