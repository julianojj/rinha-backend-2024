package domain

type Customer struct {
	Id      int64
	Name    string
	Limit   int64
	Balance int64
}

func NewCustomer(id int64) *Customer {
	return &Customer{
		Id: id,
	}
}

func (c *Customer) SetBalance(balance int64) {
	c.Balance = balance
}
