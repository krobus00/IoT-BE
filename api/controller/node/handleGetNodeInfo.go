package node

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) HandleGetNode(eCtx echo.Context) error {

	name := eCtx.Get("nodeId").(string)
	return eCtx.String(http.StatusOK, "Welcome "+name+"!")
}
