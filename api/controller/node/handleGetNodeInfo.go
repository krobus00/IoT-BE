package node

import (
	"context"
	"net/http"

	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/labstack/echo/v4"
)

func (c *Controller) HandleGetNode(eCtx echo.Context) error {

	ctx := eCtx.Request().Context()

	span := kro_util.StartTracing(ctx, tag, tracingGetNodeInfo)
	defer span.Finish()

	ctx = context.WithValue(ctx, "nodeId", eCtx.Get("nodeId").(string))
	resp, err := c.NodeService.GetNodeInfo(ctx)
	if err != nil {
		return err
	}
	response := &kro_model.Response{
		Data: resp,
	}

	return eCtx.JSON(http.StatusOK, response)
}
