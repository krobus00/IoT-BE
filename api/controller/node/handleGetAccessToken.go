package node

import (
	"net/http"

	"github.com/krobus00/iot-be/model"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/labstack/echo/v4"
)

func (c *Controller) HandleGetAccessToken(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	span := kro_util.StartTracing(ctx, tag, tracingGetAccessToken)
	defer span.Finish()

	payload := new(model.GetAccessTokenRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	}

	resp, err := c.NodeService.GetAccessToken(ctx, payload)
	if err != nil {
		return err
	}

	response := &kro_model.Response{
		Data: resp,
	}

	return eCtx.JSON(http.StatusOK, response)
}
