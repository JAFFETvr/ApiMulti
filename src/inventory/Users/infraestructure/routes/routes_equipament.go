package equipment

import (
	"database/sql"
	"log"

	controllers "gym-system/src/inventory/Users/infraestructure/controllers"
	"gym-system/src/inventory/Users/infraestructure/database"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, db *sql.DB) {
	repo, err := database.NewMySQLUserRepository()
	if err != nil {
		log.Fatalf("Error al inicializar el repositorio de usuarios: %v", err)
	}

	controller := controllers.NewUserController(repo)

	r.POST("/register", controller.RegisterUser)
	r.GET("/pending_requests", controller.GetPendingRequests)
	r.POST("/approve", controller.ApproveUser)
	r.DELETE("/pending_requests/:id", controller.RejectUser) 
	r.GET("/approved_users", controller.GetApprovedUsers)
}