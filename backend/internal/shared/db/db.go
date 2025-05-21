package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL
)

// ConnectToDB establece una conexión a la base de datos y la devuelve.
func ConnectToDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Verifica que la conexión sea válida
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la base de datos establecida correctamente.")
	return db, nil
}
