package node

import (
	"context"
	"fmt"

	db_models "github.com/krobus00/iot-be/model/database"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) Store(ctx context.Context, db kro_pkg.Querier, input *db_models.Node) error {
	span := kro_util.StartTracing(ctx, tag, tracingStore)
	defer span.Finish()

	query, args, err := r.buildInsertQuery(input).ToSql()

	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStore, err))
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStore, err))
		return err
	}

	return nil
}
