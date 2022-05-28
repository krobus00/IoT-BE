package node

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (r *repository) CountNodes(ctx context.Context, db kro_pkg.Querier) (int64, error) {
	// TODO: MOVE TO BUILDING BLOCK
	span := kro_util.StartTracing(ctx, tag, tracingCountNodes)
	defer span.Finish()

	var totalCount int64

	countQuery, args, err := sq.Select("count(id)").Where(sq.Eq{"deleted_at": nil}).From(r.GetTableName()).ToSql()
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingCountNodes, err))
		return 0, err
	}

	err = db.GetContext(ctx, &totalCount, countQuery, args...)
	if err != nil {
		r.logger.Zap.Info(fmt.Sprintf("%s %s with query: %v", tag, tracingCountNodes, countQuery))
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingCountNodes, err))
		return 0, err
	}

	return totalCount, nil
}
