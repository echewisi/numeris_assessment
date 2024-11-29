package repositories

import (
	"context"
	"github.com/echewisi/numeris_assessment/internal/models"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// CreateInvoice adds a new invoice to the database
func (r *InvoiceRepository) CreateInvoice(ctx context.Context, invoice *models.Invoice) (*models.Invoice, error) {
	if err := r.db.WithContext(ctx).Create(invoice).Error; err != nil {
		return nil, err
	}

	return invoice, nil
}

// GetInvoiceByID retrieves an invoice by its ID
func (r *InvoiceRepository) GetInvoiceByID(ctx context.Context, id int64) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.WithContext(ctx).First(&invoice, id).Error
	return &invoice, err
}

// UpdateInvoice updates an existing invoice
func (r *InvoiceRepository) UpdateInvoice(ctx context.Context, invoice *models.Invoice) (*models.Invoice, error) {
	if err :=  r.db.WithContext(ctx).Save(invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

// DeleteInvoice removes an invoice by its ID
func (r *InvoiceRepository) DeleteInvoice(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&models.Invoice{}, id).Error
}
