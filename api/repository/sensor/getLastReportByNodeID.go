package sensor

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) GetLastReportByNodeID(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) (*db_models.Sensor, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetLastReportByNodeID)
	defer span.Finish()

	result := new(db_models.Sensor)
	query, args, err := r.buildSelectQuery().
		Where(sq.Eq{
			"node_id": input.NodeID,
		}).
		OrderBy("created_at DESC").
		Limit(1).
		ToSql()

	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetLastReportByNodeID, err))
		return nil, err
	}

	err = db.GetContext(ctx, result, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetLastReportByNodeID, err))
		return nil, err
	}

	return result, nil
}
