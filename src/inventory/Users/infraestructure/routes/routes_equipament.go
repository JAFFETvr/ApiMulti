package equipment

import (
	controllers "gym-system/src/inventory/Users/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	
	controller := controllers.UserController{}

	// Registrar las rutas
	r.POST("/register", controller.RegisterUser)
	r.GET("/pending_requests", controller.GetPendingRequests)
	r.POST("/approve", controller.ApproveUser)
}
