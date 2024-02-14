package repository

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/julianojj/rinha-backend-2024/internal/core/domain"
)

type CustomerRepositoryDatabase struct {
	db *sql.DB
}

func NewCustomerRepositoryDatabase(db *sql.DB) domain.CustomerRepository {
	return &CustomerRepositoryDatabase{
		db,
	}
}

func (cr *CustomerRepositoryDatabase) FindById(id int64) (*domain.Customer, error) {
	customer := &domain.Customer{}
	err := cr.db.QueryRow("SELECT Id, Limite, Saldo FROM customers WHERE Id = $1", id).Scan(&customer.Id, &customer.Limit, &customer.Balance)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (cr *CustomerRepositoryDatabase) Update(customer *domain.Customer) error {
	_, err := cr.db.Exec("UPDATE customers SET Saldo = $1 WHERE Id = $2", customer.Balance, customer.Id)
	return err
}
