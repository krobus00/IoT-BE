package sensor

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) GetSensorByRange(ctx context.Context, db *sqlx.DB, input *model.GetProcessedDataRequest) ([]*db_models.Sensor, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetSensorByRange)
	defer span.Finish()

	results := make([]*db_models.Sensor, 0)

	query, args, err := r.buildSelectQuery().
		Where(
			sq.And{
				sq.Eq{
					"node_id": input.NodeID,
				},
				sq.GtOrEq{
					"created_at": input.StartDate,
				},
				sq.LtOrEq{
					"created_at": input.EndDate,
				},
			},
		).
		OrderBy("created_at ASC").
		ToSql()
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetSensorByRange, err))
		return nil, err
	}

	err = db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetSensorByRange, err))
		return nil, err
	}

	return results, nil
}
