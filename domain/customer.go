package domain

import "github.com/anilpdv/banking/errs"

// Customer strcut
type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// CustomerRepository : Secondary port
type CustomerRepository interface {
	FindAll(string) ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
