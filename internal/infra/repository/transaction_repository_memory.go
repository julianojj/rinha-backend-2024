package repository

import "github.com/julianojj/rinha-backend-2024/internal/core/domain"

type TransactionRepositoryMemory struct {
	Transactions []*domain.Transaction
}

func NewTransactionRepositoryMemory() domain.TransactionRepository {
	return &TransactionRepositoryMemory{
		Transactions: make([]*domain.Transaction, 0),
	}
}

func (tr *TransactionRepositoryMemory) Save(transaction *domain.Transaction) error {
	tr.Transactions = append(tr.Transactions, transaction)
	return nil
}

func (tr *TransactionRepositoryMemory) FindByCustomerId(customerId int64) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction
	for _, transaction := range tr.Transactions {
		if transaction.CustomerId == customerId {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}
