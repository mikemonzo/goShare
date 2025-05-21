package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
)

type CreateTenantInput struct {
	Name     string
	Branding string
}

type CreateTenantUseCase struct {
	tenantRepository domain.TenantRepository
}

func NewCreateTenantUseCase(tenantRepository domain.TenantRepository) *CreateTenantUseCase {
	return &CreateTenantUseCase{
		tenantRepository: tenantRepository,
	}
}

func (uc *CreateTenantUseCase) Execute(input CreateTenantInput) (*domain.Tenant, error) {
	tenant := &domain.Tenant{
		ID:        uuid.New(),
		Name:      input.Name,
		Branding:  input.Branding,
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := tenant.Validate(); err != nil {
		return nil, err
	}
	
	if err := uc.tenantRepository.Create(tenant); err != nil {
		return nil, err
	}
	return tenant, nil
}
