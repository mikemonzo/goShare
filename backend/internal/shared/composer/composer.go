package composer

import (
	"database/sql"

	"github.com/mikemonzo/goshare/internal/tenant/application"
	"github.com/mikemonzo/goshare/internal/tenant/infrastructure/adapter/inbound"
	"github.com/mikemonzo/goshare/internal/tenant/infrastructure/db"
)

// NewTenantHandler crea el controlador de Tenant con todas sus dependencias.
func NewTenantHandler(database *sql.DB) *inbound.TenantHandler {
	// Crear el repositorio
	tenantRepo := db.NewPostgresTenantRepository(database)

	// Crear el caso de uso
	createTenantUC := application.NewCreateTenantUseCase(tenantRepo)

	// Crear el controlador
	return inbound.NewTenantHandler(createTenantUC)
}
