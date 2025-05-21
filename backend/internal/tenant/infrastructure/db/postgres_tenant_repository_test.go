package db_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
	"github.com/mikemonzo/goshare/internal/tenant/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestPostgresTenantRepository_Create(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := db.NewPostgresTenantRepository(mockDB)

	tenant := &domain.Tenant{
		ID:        uuid.New(),
		Name:      "Test Tenant",
		Branding:  "Test Branding",
		IsActive:  true,
		CreatedAt: 1234567890,
		UpdatedAt: 1234567890,
	}

	mock.ExpectExec("INSERT INTO tenants").
		WithArgs(tenant.ID, tenant.Name, tenant.Branding, tenant.IsActive, tenant.CreatedAt, tenant.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(tenant)
	assert.NoError(t, err)
	mock.ExpectationsWereMet()
}
