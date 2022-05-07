package middleware

import (
	"errors"
	"net/http"
	"strings"

	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/labstack/echo/v4"
)

func DecodeJWTTokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			token := eCtx.Request().Header.Get("Authorization")
			token = strings.Replace(token, "Bearer ", "", -1)

			if token == "" {

				return kro_model.NewHttpCustomError(http.StatusUnauthorized, errors.New("Invalid Token"))
			}

			decoded, err := kro_util.DecodeTokenJWT(token, "secret")
			if err != nil {
				return kro_model.NewHttpCustomError(http.StatusUnauthorized, errors.New("Invalid Token"))
			}
			eCtx.Set("nodeId", decoded["userId"].(string))
			return next(eCtx)
		}
	}
}
