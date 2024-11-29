package repositories

import (
	"context"
	"github.com/echewisi/numeris_assessment/internal/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// CreateCustomer adds a new customer to the database
func (r *CustomerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)  {
	if err := r.db.WithContext(ctx).Create(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

// GetCustomerByID retrieves a customer by their ID
func (r *CustomerRepository) GetCustomerByID(ctx context.Context, id int64) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.WithContext(ctx).First(&customer, id).Error
	return &customer, err
}

// UpdateCustomer updates an existing customer's details
func (r *CustomerRepository) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)  {
	if err := r.db.WithContext(ctx).Save(customer).Error; err!=nil{
		return nil, err
	}
	return customer, nil
}

// DeleteCustomer removes a customer by their ID
func (r *CustomerRepository) DeleteCustomer(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&models.Customer{}, id).Error
}
