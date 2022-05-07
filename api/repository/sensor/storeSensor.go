package sensor

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) Store(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) error {
	span := kro_util.StartTracing(ctx, tag, tracingStoreSensor)
	defer span.Finish()

	query, args, err := r.buildInsertQuery(input).ToSql()

	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStoreSensor, err))
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStoreSensor, err))
		return err
	}

	return nil
}
