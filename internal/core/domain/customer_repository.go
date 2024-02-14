package domain

type CustomerRepository interface {
	FindById(id int64) (*Customer, error)
	Update(customer *Customer) error
}
