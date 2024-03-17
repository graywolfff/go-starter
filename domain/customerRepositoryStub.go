package domain

type CustomerRepositoryStub struct {
	customers []*Customer
}

func (c *CustomerRepositoryStub) FindAll() ([]*Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []*Customer{
		{
			ID:      "10001",
			Name:    "Felix",
			City:    "Hanoi",
			Zipcode: "101100",
			Status:  "active",
		},
		{
			ID:      "10002",
			Name:    "PhongDH",
			City:    "Hanoi",
			Zipcode: "101100",
			Status:  "active",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}
