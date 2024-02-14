package domain

import (
	"time"

	"github.com/julianojj/rinha-backend-2024/internal/core/exception"
)

var (
	TransactionCredit = "c"
	TransactionDebit  = "d"
)

type Transaction struct {
	CustomerId  int64
	Amount      int64
	Type        string
	Description string
	CreatedAt   time.Time
}

func NewTransaction(
	customerId int64,
	amount int64,
	transactionType string,
	description string,
) *Transaction {
	return &Transaction{
		CustomerId:  customerId,
		Amount:      amount,
		Type:        transactionType,
		Description: description,
		CreatedAt:   time.Now().UTC(),
	}
}

func (t *Transaction) Validate() error {
	if isInvalidType(t.Type) {
		return exception.NewValidationException(exception.INVALID_TRANSACTION_TYPE)
	}
	if t.Amount < 0 {
		return exception.NewValidationException(exception.INVALID_AMOUNT)
	}
	if t.Description == "" || len(t.Description) > 10 {
		return exception.NewValidationException(exception.INVALID_DESCRIPTION)
	}
	return nil
}

func isInvalidType(transactionType string) bool {
	return transactionType != TransactionCredit && transactionType != TransactionDebit
}

func (t *Transaction) GetBalance(limit, balance int64) (int64, error) {
	minBalance := limit * -1
	if t.IsDebit() {
		balance -= t.Amount
		if balance < minBalance {
			return 0, exception.NewValidationException(exception.INSUFFICIENT_FUNDS)
		}
	} else {
		balance += t.Amount
	}
	return balance, nil
}

func (t *Transaction) IsDebit() bool {
	return t.Type == TransactionDebit
}
