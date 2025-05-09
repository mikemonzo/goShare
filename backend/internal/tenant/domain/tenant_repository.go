package domain

import "github.com/google/uuid"

type TenantRepository interface {
	Create(tenant *Tenant) error
	GetByID(id uuid.UUID) (*Tenant, error)
	GetByName(name string) (*Tenant, error)
	Update(tenant *Tenant) error
	Delete(id uuid.UUID) error
	ListActive() ([]*Tenant, error)
	GetAll() ([]*Tenant, error)
}
