package expense

import "database/sql"

type ExpenseRepository struct {
	Db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{
		Db: db,
	}
}

func (e *ExpenseRepository) Add(id string, request *NewExpenseRequest) error {
	query := "INSERT INTO expense(id, item, quantity, price, wallet_id) VALUES ($1, $2, $3, $4,$5)"
	_, err := e.Db.Exec(query, id, request.Item, request.Quantity, request.Price, request.WalletId)

	return err
}

func (e *ExpenseRepository) FindById(id string) (*Expense, error) {
	query := "SELECT id, item, quantity, price, total_price, wallet_id, created_at FROM expense WHERE id=$1"
	result := e.Db.QueryRow(query, id)

	if result.Err() != nil {
		return nil, result.Err()
	}

	data := new(Expense)
	if err := result.Scan(&data.Id, &data.Item, &data.Quantity, &data.Price, &data.TotalPrice, &data.WalletId, &data.CreatedAt); err != nil {
		return nil, err
	}

	return data, nil
}

func (e *ExpenseRepository) FindAll() (*[]Expense, error) {
	query := "SELECT id, item, quantity, price, total_price, wallet_id, created_at FROM expense"
	result, err := e.Db.Query(query)
	if err != nil {
		return nil, err
	}

	list := []Expense{}
	for result.Next() {
		data := Expense{}
		err := result.Scan(&data.Id, &data.Item, &data.Quantity, &data.Price, &data.TotalPrice, &data.WalletId, &data.CreatedAt)
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

func (e *ExpenseRepository) RemoveById(id string) error {
	query := "DELETE FROM expense WHERE id=$1"
	_, err := e.Db.Exec(query, id)

	return err
}

func (e *ExpenseRepository) FindAllByWalletId(idWallet string) (*[]Expense, error) {
	query := "SELECT id, item, quantity, price, total_price, wallet_id, created_at FROM expense WHERE wallet_id=$1"
	result, err := e.Db.Query(query, idWallet)
	if err != nil {
		return nil, err
	}

	list := []Expense{}
	for result.Next() {
		data := Expense{}
		err := result.Scan(&data.Id, &data.Item, &data.Quantity, &data.Price, &data.TotalPrice, &data.WalletId, &data.CreatedAt)
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
