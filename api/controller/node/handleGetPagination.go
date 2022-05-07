package node

import (
	"net/http"

	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/labstack/echo/v4"
)

func (c *Controller) HandleGetPagination(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	span := kro_util.StartTracing(ctx, tag, tracingGetPagination)
	defer span.Finish()

	payload := new(kro_model.PaginationRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	}

	resp, err := c.NodeService.GetAllNodes(ctx, payload)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}
