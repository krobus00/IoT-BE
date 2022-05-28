package node

import (
	"fmt"
	"net/http"

	"github.com/krobus00/iot-be/model"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/labstack/echo/v4"
)

func (c *Controller) HandleRegister(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	span := kro_util.StartTracing(ctx, tag, tracingRegister)
	defer span.Finish()

	payload := new(model.RegisterRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	}

	id, err := c.NodeService.Register(ctx, payload)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("new node with id %s", *id)
	resp := kro_model.Response{
		Message: &message,
	}

	return eCtx.JSON(http.StatusOK, resp)
}
