package node

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) GetAllNodes(ctx context.Context, db *sqlx.DB, paginationRequest *kro_model.PaginationRequest, config ...kro_model.Config) ([]*db_models.Node, int64, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetAllNodes)
	defer span.Finish()

	var totalCount int64
	results := make([]*db_models.Node, 0)

	paginationConfig := kro_model.ParamsFromRequest(paginationRequest)

	paginationSearchField := kro_model.SearchFields(searchFields...)
	selectBuilder, countBuilder, err := kro_model.PaginateQuery(ctx, db, results, r.buildSelectQuery(), paginationConfig, columnMapping, whiteList, paginationSearchField, config)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllNodes, err))
		return nil, 0, err
	}

	countQuery, args, err := countBuilder.ToSql()
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllNodes, err))
		return nil, 0, err
	}
	selectQuery, args, err := selectBuilder.ToSql()
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllNodes, err))
		return nil, 0, err
	}

	err = db.GetContext(ctx, &totalCount, countQuery, args...)
	if err != nil {
		r.logger.Zap.Info(fmt.Sprintf("%s %s with query: %v", tag, tracingGetAllNodes, countQuery))
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllNodes, err))
		return nil, 0, err
	}
	err = db.SelectContext(ctx, &results, selectQuery, args...)
	if err != nil {
		r.logger.Zap.Info(fmt.Sprintf("%s %s with query: %v", tag, tracingGetAllNodes, selectQuery))
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllNodes, err))
		return nil, 0, err
	}
	r.logger.Zap.Info(fmt.Sprintf("%s %s with query: %v", tag, tracingGetAllNodes, selectQuery))
	return results, totalCount, nil
}
