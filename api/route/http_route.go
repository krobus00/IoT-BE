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

	internal := e.Group("/_internal/iot")

	internal.POST("/node/register", handler.NodeController.HandleRegister)
	internal.POST("/node/getToken", handler.NodeController.HandleGetAccessToken)

	iot := e.Group("/iot")

	iot.GET("/node", handler.NodeController.HandleGetPagination)
	iot.GET("/node/me", handler.NodeController.HandleGetNode, local_middleware.DecodeJWTTokenMiddleware())

	iot.GET("/sensors", handler.SensorController.HandleGetAllSensor)
	iot.POST("/sensors", handler.SensorController.HandleCreateSensor, local_middleware.DecodeJWTTokenMiddleware())
	iot.POST("/sensors/resampledData/:nodeId", handler.SensorController.HandleGetResampledData)
	iot.GET("/sensors/:id", handler.SensorController.HandleShowSensor)
	iot.PATCH("/sensors/:id", handler.SensorController.HandleUpdateSensor)
	iot.DELETE("/sensors/:id", handler.SensorController.HandleDeleteSensor)

	// INJECT ROUTE HERE
}

var Module = fx.Options(
	fx.Invoke(NewRoutes),
)
