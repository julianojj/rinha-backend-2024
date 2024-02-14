package repository

import (
	"database/sql"

	"github.com/julianojj/rinha-backend-2024/internal/core/domain"
)

type TransactionRepositoryDatabase struct {
	db *sql.DB
}

func NewTransactionRepositoryDatabase(db *sql.DB) domain.TransactionRepository {
	return &TransactionRepositoryDatabase{
		db,
	}
}

func (tr *TransactionRepositoryDatabase) Save(transaction *domain.Transaction) error {
	_, err := tr.db.Exec("INSERT INTO transactions(ClienteID, Tipo, Descricao, Valor, DataCriacao) VALUES($1, $2, $3, $4, $5)",
		transaction.CustomerId, transaction.Type, transaction.Description, transaction.Amount, transaction.CreatedAt)
	return err
}

func (tr *TransactionRepositoryDatabase) FindByCustomerId(customerId int64) ([]*domain.Transaction, error) {
	rows, err := tr.db.Query("SELECT Tipo, Descricao, Valor, DataCriacao FROM transactions WHERE ClienteID = $1 ORDER BY DataCriacao DESC LIMIT 10", customerId)
	if err != nil {
		return nil, err
	}
	var transactions []*domain.Transaction
	for rows.Next() {
		var transaction domain.Transaction
		if err := rows.Scan(&transaction.Type, &transaction.Description, &transaction.Amount, &transaction.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}
