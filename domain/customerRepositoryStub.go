package domain

// CustomerRepositoryStub : struct
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll : recevier
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

// NewCustomerRepositoryStub : stub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "jfdkfjd", Name: "arnodh", City: "hyderabad", Zipcode: "kjfdkjkf"},
		{ID: "dkfjdk", Name: "ekjr", City: "hyderabad", Zipcode: "kjfdkjkf"},
	}

	return CustomerRepositoryStub{customers: customers}
}
