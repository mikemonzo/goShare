package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        uuid.UUID
	Name      string
	IsActive  bool
	Branding  string
	CreatedAt int64
	UpdatedAt int64
}

func (t *Tenant) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
