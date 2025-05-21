package app

import (
	"fmt"
	"log"

	"github.com/mikemonzo/goshare/internal/shared/composer"
	"github.com/mikemonzo/goshare/internal/shared/db"
	"github.com/mikemonzo/goshare/internal/shared/logger"
	"github.com/mikemonzo/goshare/internal/shared/router"
	"github.com/mikemonzo/goshare/pkg/config"
)

func main() {
	logger.InitLogger()
	// Load .env
	cfg := config.LoadConfig()

	// TODO: Initialize routes, servicies, DB, middlewares, etc.

	// Initialize DB connection
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	database, err := db.ConnectToDB(dbConnectionString)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer database.Close()

	logger.Log.Infof("Servidor escuchando en el puerto %s...", cfg.AppPort)

	// Usar el compositor para crear el controlador
	tenantHandler := composer.NewTenantHandler(database)

	// Configurar rutas
	rout := router.SetRouter(tenantHandler)

	logger.Log.Infof("Servidor escuchando en el puerto %s...", cfg.AppPort)

	if err := rout.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
