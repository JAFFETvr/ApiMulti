package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	equipamentInfra "gym-system/src/inventory/Users/infraestructure/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Conectar a la base de datos
	db, err := sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	// Configurar servidor Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Registrar rutas con la base de datos
	equipamentInfra.RegisterUserRoutes(r, db)

	// Definir puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
