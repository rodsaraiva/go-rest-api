package service

import "github.com/rodsaraiva/go-rest-api/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaltCustomerService struct {
	rep domain.CustomerRepository
}

func (s DefaltCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.rep.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaltCustomerService {
	return DefaltCustomerService{rep: repository}
}
