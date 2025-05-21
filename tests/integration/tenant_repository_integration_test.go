package integration

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
	"github.com/stretchr/testify/assert"
)

func setupDatabase(t *testing.T) *sql.DB {
	// Configurar conexión a la base de datos
	dsn := "postgres://postgres:postgres@localhost:5432/goshare?sslmode=disable"
	database, err := sql.Open("postgres", dsn)
	assert.NoError(t, err)

	// Verificar conexión
	err = database.Ping()
	assert.NoError(t, err)

	// Limpiar tabla antes de cada prueba
	_, err = database.Exec("DELETE FROM tenants")
	assert.NoError(t, err)

	return database
}

func TestPostgresTenantRepository_CreateAndGetByID(t *testing.T) {
	db := setupDatabase(t)
	defer db.Close()

	repo := db.NewPostgresTenantRepository(db)

	// Crear un tenant
	tenant := &domain.Tenant{
		ID:        uuid.New(),
		Name:      "Test Tenant",
		Branding:  "Test Branding",
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := repo.Create(tenant)
	assert.NoError(t, err)

	// Obtener el tenant por ID
	retrievedTenant, err := repo.GetByID(tenant.ID)
	assert.NoError(t, err)
	assert.Equal(t, tenant.Name, retrievedTenant.Name)
	assert.Equal(t, tenant.Branding, retrievedTenant.Branding)
	assert.Equal(t, tenant.IsActive, retrievedTenant.IsActive)
}

func TestPostgresTenantRepository_GetByName(t *testing.T) {
	db := setupDatabase(t)
	defer db.Close()

	repo := db.NewPostgresTenantRepository(db)

	// Crear un tenant
	tenant := &domain.Tenant{
		ID:        uuid.New(),
		Name:      "Test Tenant",
		Branding:  "Test Branding",
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := repo.Create(tenant)
	assert.NoError(t, err)

	// Obtener el tenant por nombre
	retrievedTenant, err := repo.GetByName(tenant.Name)
	assert.NoError(t, err)
	assert.Equal(t, tenant.ID, retrievedTenant.ID)
	assert.Equal(t, tenant.Name, retrievedTenant.Name)
	assert.Equal(t, tenant.Branding, retrievedTenant.Branding)
}

func TestPostgresTenantRepository_Delete(t *testing.T) {
	db := setupDatabase(t)
	defer db.Close()

	repo := db.NewPostgresTenantRepository(db)

	// Crear un tenant
	tenant := &domain.Tenant{
		ID:        uuid.New(),
		Name:      "Test Tenant",
		Branding:  "Test Branding",
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := repo.Create(tenant)
	assert.NoError(t, err)

	// Eliminar el tenant
	err = repo.Delete(tenant.ID)
	assert.NoError(t, err)

	// Intentar obtener el tenant eliminado
	retrievedTenant, err := repo.GetByID(tenant.ID)
	assert.Error(t, err)
	assert.Nil(t, retrievedTenant)
}
