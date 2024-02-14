package repository

import "github.com/julianojj/rinha-backend-2024/internal/core/domain"

type CustomerRepositoryMemory struct {
	Customers []*domain.Customer
}

func NewCustomerRepository() domain.CustomerRepository {
	customers := []*domain.Customer{
		{
			Id:      1,
			Limit:   100000,
			Balance: 0,
		},
		{
			Id:      2,
			Limit:   80000,
			Balance: 0,
		},
		{
			Id:      3,
			Limit:   1000000,
			Balance: 0,
		},
		{
			Id:      4,
			Limit:   10000000,
			Balance: 0,
		},
		{
			Id:      5,
			Limit:   500000,
			Balance: 0,
		},
	}
	return &CustomerRepositoryMemory{
		Customers: customers,
	}
}

func (cr *CustomerRepositoryMemory) FindById(id int64) (*domain.Customer, error) {
	for _, customer := range cr.Customers {
		if customer.Id == id {
			return customer, nil
		}
	}
	return nil, nil
}

func (cr *CustomerRepositoryMemory) Update(customer *domain.Customer) error {
	for _, c := range cr.Customers {
		if c == customer {
			c = customer
		}
	}
	return nil
}
