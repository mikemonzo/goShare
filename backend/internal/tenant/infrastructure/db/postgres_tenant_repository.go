package db

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
)

type PostgresTenantRepository struct {
	DB *sql.DB
}

func NewPostgresTenantRepository(db *sql.DB) *PostgresTenantRepository {
	return &PostgresTenantRepository{
		DB: db,
	}
}

func (r *PostgresTenantRepository) Create(tenant *domain.Tenant) error {
	query := `INSERT INTO tenants (id, name, branding, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.DB.Exec(query, tenant.ID, tenant.Name, tenant.Branding, tenant.IsActive, tenant.CreatedAt, tenant.UpdatedAt)

	return err
}

func (r *PostgresTenantRepository) GetByID(id uuid.UUID) (*domain.Tenant, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *PostgresTenantRepository) GetByName(name string) (*domain.Tenant, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *PostgresTenantRepository) Update(tenant *domain.Tenant) error {
	// TODO: Implement this method
	return nil
}

func (r *PostgresTenantRepository) Delete(id uuid.UUID) error {
	// TODO: Implement this method
	return nil
}

func (r *PostgresTenantRepository) ListActive() ([]*domain.Tenant, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *PostgresTenantRepository) GetAll() ([]*domain.Tenant, error) {
	// TODO: Implement this method
	return nil, nil
}
