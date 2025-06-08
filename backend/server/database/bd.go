package database

import (
	"database/sql"
	"fmt"
	"log"

	"app-profesor/backend/server/database/scripts"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=adiel password=Adiel123 dbname=atcnea sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error de conexión:", err)
	}

	// Verificar conexión
	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a PostgreSQL:", err)
	}
	fmt.Println("Conexión exitosa a PostgreSQL")

	scripts.EjecuteScripts(DB)

}
