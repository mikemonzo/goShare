package domain

import "github.com/google/uuid"

type Tenant struct {
	ID        uuid.UUID
	Name      string
	IsActive  bool
	Branding  string
	CreatedAt int64
	UpdatedAt int64
}
