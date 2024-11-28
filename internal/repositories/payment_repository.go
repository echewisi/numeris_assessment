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

// CreatePayment adds a new payment to the database
func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

// GetPaymentByID retrieves a payment by its ID
func (r *PaymentRepository) GetPaymentByID(ctx context.Context, id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.WithContext(ctx).First(&payment, id).Error
	return &payment, err
}

// UpdatePayment updates an existing payment
func (r *PaymentRepository) UpdatePayment(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

// DeletePayment removes a payment by its ID
func (r *PaymentRepository) DeletePayment(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Payment{}, id).Error
}
