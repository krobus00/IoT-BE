package sensor

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) GetSensorByID(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) (*db_models.Sensor, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetSensorByID)
	defer span.Finish()

	result := new(db_models.Sensor)
	query, args, err := r.buildSelectQuery().Where(sq.Eq{
		"id": input.ID,
	}).ToSql()

	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetSensorByID, err))
		return nil, err
	}

	err = db.GetContext(ctx, result, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetSensorByID, err))
		return nil, err
	}

	return result, nil
}
