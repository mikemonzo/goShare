package domain

import "github.com/google/uuid"

type UserRepository interface {
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(tenantrID uuid.UUID, email string) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
	ListByTenant(tenantID uuid.UUID) ([]*User, error)
}
