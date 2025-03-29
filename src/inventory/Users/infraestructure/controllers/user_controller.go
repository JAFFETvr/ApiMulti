package equipamentcontrollers

import (

	"github.com/gin-gonic/gin"
	equipmentusecases "gym-system/src/inventory/Users/application/useCases"
	"gym-system/src/inventory/Users/domain/repository"
	"net/http"
	"strconv"
)

type UserController struct {
	Repo repository.IUserRepository
}

func NewUserController(repo repository.IUserRepository) *UserController {
	return &UserController{Repo: repo}
}

// Adaptar la función RegisterUser a Gin
func (c *UserController) RegisterUser(ctx *gin.Context) {
	var request map[string]string
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// Si hay error al leer el JSON, respondemos con error 400
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Llamada al repositorio para guardar la solicitud
	err := c.Repo.SaveRequest(request["username"], request["password"])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respondemos con código 201 (Created) si todo salió bien
	ctx.JSON(http.StatusCreated, gin.H{"message": "Solicitud enviada"})
}

// Adaptar la función GetPendingRequests a Gin
func (c *UserController) GetPendingRequests(ctx *gin.Context) {
	// Obtener las solicitudes pendientes desde el repositorio
	requests, err := c.Repo.GetPendingRequests()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enviar la respuesta en formato JSON
	ctx.JSON(http.StatusOK, requests)
}

// Adaptar la función ApproveUser a Gin
func (c *UserController) ApproveUser(ctx *gin.Context) {
	// Obtener parámetros de la URL
	id, _ := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	macAddress := ctx.DefaultQuery("mac", "")

	// Crear el caso de uso para aprobar al usuario
	useCase := equipmentusecases.NewApproveUserUseCase(c.Repo)
	err := useCase.Execute(id, macAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con éxito
	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario aprobado"})
}
