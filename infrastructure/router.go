package infrastructure

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo *echo.Echo
}

func NewRouter(env Env, customValidator echo.Validator) Router {

	e := kro_pkg.NewRouter()

	e.Validator = customValidator

	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	return Router{
		Echo: e,
	}
}
