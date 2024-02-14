package services

import (
	"time"

	"github.com/julianojj/rinha-backend-2024/internal/core/domain"
	"github.com/julianojj/rinha-backend-2024/internal/core/exception"
)

type (
	TransactionService struct {
		customerRepository    domain.CustomerRepository
		transactionRepository domain.TransactionRepository
	}
	ProcessTransactionInput struct {
		CustomerId  int64  `json:"customer_id,omitempty"`
		Amount      int64  `json:"valor"`
		Type        string `json:"tipo"`
		Description string `json:"descricao"`
	}
	ProcessTransactionOutput struct {
		Limit   int64 `json:"limite"`
		Balance int64 `json:"saldo"`
	}
	RequestStatementOutput struct {
		Balance      *RequestStatementBalanceOutput        `json:"saldo"`
		Transactions []*RequestStatementTransactionsOuptut `json:"ultimas_transacoes"`
	}
	RequestStatementBalanceOutput struct {
		Total     int64     `json:"total"`
		IssueDate time.Time `json:"data_extrato"`
		Limit     int64     `json:"limite"`
	}
	RequestStatementTransactionsOuptut struct {
		Amount      int64     `json:"valor"`
		Type        string    `json:"tipo"`
		Description string    `json:"descricao"`
		CreatedAt   time.Time `json:"realizada_em"`
	}
)

func NewTransactionService(
	customerRepository domain.CustomerRepository,
	transactionRepository domain.TransactionRepository,
) *TransactionService {
	return &TransactionService{
		customerRepository,
		transactionRepository,
	}
}

func (ts *TransactionService) ProcessTransaction(input *ProcessTransactionInput) (*ProcessTransactionOutput, error) {
	existingCustomer, err := ts.customerRepository.FindById(input.CustomerId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if existingCustomer == nil {
		return nil, exception.NewNotFoundException(exception.CUSTOMER_NOT_FOUND)
	}
	transaction := domain.NewTransaction(
		existingCustomer.Id,
		input.Amount,
		input.Type,
		input.Description,
	)
	if err := transaction.Validate(); err != nil {
		return nil, err
	}
	balance, err := transaction.GetBalance(existingCustomer.Limit, existingCustomer.Balance)
	if err != nil {
		return nil, err
	}
	existingCustomer.Balance = balance
	if err := ts.transactionRepository.Save(transaction); err != nil {
		return nil, err
	}
	ts.customerRepository.Update(existingCustomer)
	return &ProcessTransactionOutput{
		Limit:   existingCustomer.Limit,
		Balance: balance,
	}, nil
}

func (ts *TransactionService) RequestStatement(customerId int64) (*RequestStatementOutput, error) {
	existingCustomer, err := ts.customerRepository.FindById(customerId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if existingCustomer == nil {
		return nil, exception.NewNotFoundException(exception.CUSTOMER_NOT_FOUND)
	}
	transactions, err := ts.transactionRepository.FindByCustomerId(existingCustomer.Id)
	if err != nil {
		return nil, err
	}
	var transactionsOuptput []*RequestStatementTransactionsOuptut
	for _, transaction := range transactions {
		transactionsOuptput = append(transactionsOuptput, &RequestStatementTransactionsOuptut{
			Amount:      transaction.Amount,
			Type:        transaction.Type,
			Description: transaction.Description,
			CreatedAt:   transaction.CreatedAt,
		})
	}
	output := &RequestStatementOutput{
		Balance: &RequestStatementBalanceOutput{
			Total:     existingCustomer.Balance,
			IssueDate: time.Now().UTC(),
			Limit:     existingCustomer.Limit,
		},
		Transactions: transactionsOuptput,
	}
	return output, nil
}
