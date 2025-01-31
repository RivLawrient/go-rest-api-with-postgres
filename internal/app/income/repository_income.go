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
