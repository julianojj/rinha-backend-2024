package domain

type TransactionRepository interface {
	Save(transaction *Transaction) error
	FindByCustomerId(customerId int64) ([]*Transaction, error)
}
