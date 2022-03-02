package service

import (
	"github.com/rodsaraiva/go-rest-api/domain"
	"github.com/rodsaraiva/go-rest-api/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaltCustomerService struct {
	rep domain.CustomerRepository
}

func (s DefaltCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.rep.FindAll()
}

func (s DefaltCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.rep.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaltCustomerService {
	return DefaltCustomerService{rep: repository}
}
