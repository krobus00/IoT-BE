package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/krobus00/iot-be/api/controller"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/api/route"
	"github.com/krobus00/iot-be/api/service"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/requester"
	"go.uber.org/fx"
)

var AppModule = fx.Options(
	infrastructure.Module,
	requester.Module,
	repository.Module,
	service.Module,
	controller.Module,
	route.Module,
	fx.Invoke(appBootstrap),
)

func appBootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
) {

	appStop := func(context.Context) error {
		logger.Zap.Info("Stopping Application")
		conn := database.DB
		conn.Close()
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			figure.NewColorFigure(env.AppName, "", "purple", true).Print()
			go func() {
				err := database.DB.Ping()
				if err != nil {
					logger.Zap.Panic(err)
				} else {
					logger.Zap.Info("Database connected")

				}
				PORT := "5000"
				if env.AppPort != "" {
					PORT = env.AppPort
				}
				logger.Zap.Info(fmt.Sprintf("APP RUNNING ON http://0.0.0.0:%s", PORT))
				if err := handler.Echo.Start(fmt.Sprintf(":%s", PORT)); err != nil && err != http.ErrServerClosed {
					handler.Echo.Logger.Fatal("shutting down the server")
				}
			}()
			return nil
		},
		OnStop: appStop,
	})
}
