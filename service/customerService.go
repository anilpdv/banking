package service

import (
	"github.com/anilpdv/banking/domain"
	"github.com/anilpdv/banking/errs"
)

// CustomerService : interface
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService : struct
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer : receiver
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

// GetCustomer : receiver
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// NewCustomerService : func
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
