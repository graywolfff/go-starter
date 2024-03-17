package service

import "github.com/graywolff/go-banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (r *DefaultCustomerService) GetAllCustomer() ([]*domain.Customer, error) {
	return r.repo.FindAll()
}
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
