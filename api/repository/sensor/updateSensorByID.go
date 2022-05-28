package sensor

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) UpdateSensorByID(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) error {
	span := kro_util.StartTracing(ctx, tag, tracingUpdateSensorByID)
	defer span.Finish()

	query, args, err := r.buildUpdateQuery(input).Where(sq.Eq{
		"id": input.ID,
	}).ToSql()

	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingUpdateSensorByID, err))
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingUpdateSensorByID, err))
		return err
	}

	return nil
}
