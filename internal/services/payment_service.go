package services

import (
	"context"
	"errors"

	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/repositories"
)

type PaymentService struct {
	Repo *repositories.PaymentRepository
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{Repo: repo}
}

// CreatePayment processes a new payment
func (s *PaymentService) CreatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	if payment.Amount <= 0 {
		return nil, errors.New("payment amount must be greater than zero")
	}

	return s.Repo.CreatePayment(ctx, payment)
}

// GetPayment retrieves a payment by ID
func (s *PaymentService) GetPayment(ctx context.Context, id int64) (*models.Payment, error) {
	return s.Repo.GetPaymentByID(ctx, id)
}

// UpdatePayment updates an existing payment
func (s *PaymentService) UpdatePayment(ctx context.Context, updatedPayment *models.Payment) (*models.Payment, error) {
	return s.Repo.UpdatePayment(ctx, updatedPayment)
}

// DeletePayment deletes a payment by ID
func (s *PaymentService) DeletePayment(ctx context.Context, id int64) error {
	return s.Repo.DeletePayment(ctx, id)
}
