package domain

import "github.com/google/uuid"

type Role string

const (
	RoleSystemAdmin Role = "SYSTEM_ADMIN"
	RoleTenantAdmin Role = "TENANT_ADMIN"
	RoleUser        Role = "USER"
	RoleGuest       Role = "GUEST"
)

type User struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      Role
	IsActive  bool
	CreatedAt int64
	UpdatedAt int64
}
