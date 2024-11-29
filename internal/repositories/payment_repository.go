package repositories

import (
	"context"
	"github.com/echewisi/numeris_assessment/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// CreatePayment adds a new payment to the database and returns the created payment
func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	if err := r.db.WithContext(ctx).Create(payment).Error; err != nil {
		return nil, err
	}
	// Return the created payment
	return payment, nil
}


// GetPaymentByID retrieves a payment by its ID
func (r *PaymentRepository) GetPaymentByID(ctx context.Context, id int64) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.WithContext(ctx).First(&payment, id).Error
	return &payment, err
}

// UpdatePayment updates an existing payment
func (r *PaymentRepository) UpdatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	if err := r.db.WithContext(ctx).Save(payment).Error; err != nil{
		return nil, err
	}

	return payment, nil
}

// DeletePayment removes a payment by its ID
func (r *PaymentRepository) DeletePayment(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&models.Payment{}, id).Error
}
