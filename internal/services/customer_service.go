package services

import (
	"context"
	"errors"

	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/repositories"
)

type CustomerService struct {
	Repo *repositories.CustomerRepository
}

// NewCustomerService creates a new CustomerService
func NewCustomerService(repo *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{Repo: repo}
}

// CreateCustomer creates a new customer
func (s *CustomerService) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if customer.Email == "" {
		return nil, errors.New("email is required")
	}

	return s.Repo.CreateCustomer(ctx, customer)
}

// GetCustomer retrieves a customer by ID
func (s *CustomerService) GetCustomer(ctx context.Context, id int64) (*models.Customer, error) {
	return s.Repo.GetCustomerByID(ctx, id)
}

// UpdateCustomer updates an existing customer
func (s *CustomerService) UpdateCustomer(ctx context.Context, updatedCustomer *models.Customer) (*models.Customer, error) {
	return s.Repo.UpdateCustomer(ctx, updatedCustomer)
}

// DeleteCustomer deletes a customer by ID
func (s *CustomerService) DeleteCustomer(ctx context.Context, id int64) error {
	return s.Repo.DeleteCustomer(ctx, id)
}
