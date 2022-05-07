package sensor

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) DeleteSensorByID(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) error {
	span := kro_util.StartTracing(ctx, tag, tracingDeleteSensorByID)
	defer span.Finish()

	query, args, err := r.buildDeleteQuery().Where(sq.Eq{
		"id": input.ID,
	}).ToSql()

	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingDeleteSensorByID, err))
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingDeleteSensorByID, err))
		return err
	}

	return nil
}
