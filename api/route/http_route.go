package route

import (
	"github.com/krobus00/iot-be/api/controller"
	local_middleware "github.com/krobus00/iot-be/api/middleware"
	"github.com/krobus00/iot-be/infrastructure"
	"go.uber.org/fx"
)

func NewRoutes(
	router infrastructure.Router,
	handler controller.Handler,
) {
	e := router.Echo

	internal := e.Group("/_internal")

	internal.POST("/node/register", handler.NodeController.HandleRegister)
	internal.POST("/node/getToken", handler.NodeController.HandleGetAccessToken)

	e.GET("/node", handler.NodeController.HandleGetPagination)
	e.GET("/node/me", handler.NodeController.HandleGetNode, local_middleware.DecodeJWTTokenMiddleware())

	e.GET("/sensors", handler.SensorController.HandleGetAllSensor)
	e.POST("/sensors", handler.SensorController.HandleCreateSensor, local_middleware.DecodeJWTTokenMiddleware())
	e.GET("/sensors/:id", handler.SensorController.HandleShowSensor)
	e.PATCH("/sensors/:id", handler.SensorController.HandleUpdateSensor)
	e.DELETE("/sensors/:id", handler.SensorController.HandleDeleteSensor)

	// INJECT ROUTE HERE
}

var Module = fx.Options(
	fx.Invoke(NewRoutes),
)
