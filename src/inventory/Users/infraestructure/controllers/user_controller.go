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

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var request map[string]string
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := c.Repo.SaveRequest(request["username"], request["password_hash"])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Solicitud enviada"})
}
func (c *UserController) GetPendingRequests(ctx *gin.Context) {
	requests, err := c.Repo.GetPendingRequests()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, requests)
}

func (c *UserController) ApproveUser(ctx *gin.Context) {
	// Obtener parámetros de la URL
	id, _ := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	macAddress := ctx.DefaultQuery("mac", "")

	useCase := equipmentusecases.NewApproveUserUseCase(c.Repo)
	err := useCase.Execute(id, macAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario aprobado"})
}

func (c *UserController) RejectUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := equipmentusecases.NewRejectUserUseCase(c.Repo)
	err = useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Solicitud de usuario rechazada"})
}

func (c *UserController) GetApprovedUsers(ctx *gin.Context) {
	useCase := equipmentusecases.NewGetApprovedUsersUseCase(c.Repo)
	users, err := useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}