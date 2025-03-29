package machine

import (

	machineControllers "gym-system/src/inventory/Machines/infraestructure/controllers"
	"gym-system/src/inventory/Machines/infraestructure/database"
	"gym-system/src/inventory/Machines/application/useCases"

	"github.com/gin-gonic/gin"
)

func SetupRoutesMachine(r *gin.Engine){

	dbInstance := machine.NewMySQLMachine()

	listMachineController := machineControllers.NewListMachineController(*machineusecases.NewListMachine(dbInstance))
	createMachineController := machineControllers.NewCreateMachineController(*machineusecases.NewCreateMachine(dbInstance))
	getMachineById := machineControllers.NewMachineByIdController(*machineusecases.NewMachineById(dbInstance))
	getStatusMachine := machineControllers.NewStatusMachine(machineusecases.NewMachineStatus(dbInstance))
	updateMachine := machineControllers.NewUpdateMachineController(*machineusecases.NewUpdateMachine(dbInstance))
	deleteMachine := machineControllers.NewDeleteMachine(*machineusecases.NewDeleteMachine(dbInstance))

	r.GET("/machines",listMachineController.Execute)
	r.POST("/machines",createMachineController.Execute)
	r.GET("/machines/:id",getMachineById.Execute)
	r.GET("/machines/status/:id",getStatusMachine.Execute)
	r.PUT("/machines/:id",updateMachine.Execute)
	r.DELETE("/machines/:id",deleteMachine.Execute)

}