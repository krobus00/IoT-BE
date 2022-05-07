package sensor

import (
	"net/http"

	"github.com/krobus00/iot-be/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/labstack/echo/v4"
)

func (c *Controller) HandleShowSensor(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()
	span := kro_util.StartTracing(ctx, tag, tracingShowSensor)
	defer span.Finish()

	payload := new(model.ShowSensorRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	}

	resp, err := c.SensorService.GetSensorByID(ctx, payload)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}
