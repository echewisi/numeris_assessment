package services

import (
	"context"
	"time"

	"github.com/echewisi/numeris_assessment/internal/models"
	"github.com/echewisi/numeris_assessment/internal/repositories"
)

type InvoiceService struct {
	Repo *repositories.InvoiceRepository
}

// NewInvoiceService creates a new InvoiceService
func NewInvoiceService(repo *repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{Repo: repo}
}

// CreateInvoice creates a new invoice
func (s *InvoiceService) CreateInvoice(ctx context.Context, invoice *models.Invoice) (*models.Invoice, error) {
	// if invoice.Amount <= 0 {
	// 	return errors.New("amount must be greater than zero")
	// }

	// Set default values if needed
	if invoice.Status == "" {
		invoice.Status = "unpaid"
	}
	if invoice.CreatedAt.IsZero() {
		invoice.CreatedAt = time.Now()
	}

	return s.Repo.CreateInvoice(ctx, invoice)
}

// GetInvoiceByID retrieves an invoice by its ID
func (s *InvoiceService) GetInvoiceByID(ctx context.Context, id int64) (*models.Invoice, error) {
	invoice, err := s.Repo.GetInvoiceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// UpdateInvoice updates an existing invoice
func (s *InvoiceService) UpdateInvoice(ctx context.Context, invoice *models.Invoice) (*models.Invoice, error) {
	// if invoice.ID == 0 {
	// 	return errors.New("invoice ID is required for update")
	// }

	return s.Repo.UpdateInvoice(ctx, invoice)
}

// DeleteInvoice removes an invoice by its ID
func (s *InvoiceService) DeleteInvoice(ctx context.Context, id int64) error {
	return s.Repo.DeleteInvoice(ctx, id)
}
