package sensor

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

var (
	searchFields = []string{"id", "node_id"}
	whiteList    = kro_model.ColumnWhiteList{
		"id":          true,
		"node_id":     true,
		"humidity":    true,
		"temperature": true,
		"heat_index":  true,
		"created_at":  true,
		"updated_at":  true,
		"deleted_at":  true,
	}
	columnMapping = kro_model.ColumnMapping{
		"id":          "id",
		"nodeId":      "node_id",
		"humidity":    "humidity",
		"temperature": "temperature",
		"heatIndex":   "heat_index",
		"createdAt":   "created_at",
		"updatedAt":   "updated_at",
		"deletedAt":   "deleted_at",
	}
)

func (r *repository) buildInsertQuery(input *db_models.Sensor) sq.InsertBuilder {
	vals := sq.Eq{
		"id":          input.ID,
		"node_id":     input.NodeID,
		"humidity":    input.Humidity,
		"temperature": input.Temperature,
		"heat_index":  input.HeatIndex,
	}
	insertBuilder := sq.Insert(r.GetTableName()).SetMap(vals)
	return insertBuilder
}

func (r *repository) buildSelectQuery() sq.SelectBuilder {
	selection := []string{
		"id",
		"node_id",
		"humidity",
		"temperature",
		"heat_index",
		"created_at",
		"updated_at",
		"deleted_at",
	}
	selectBuilder := sq.Select(selection...).Where(sq.Eq{"deleted_at": nil}).From(r.GetTableName())
	return selectBuilder
}

func (r *repository) buildUpdateQuery(input *db_models.Sensor) sq.UpdateBuilder {
	vals := sq.Eq{
		"node_id":     input.NodeID,
		"humidity":    input.Humidity,
		"temperature": input.Temperature,
		"heat_index":  input.HeatIndex,
		"updated_at":  time.Now().Unix(),
	}
	updateBuilder := sq.Update(r.GetTableName()).SetMap(vals)
	return updateBuilder
}

func (r *repository) buildDeleteQuery() sq.UpdateBuilder {
	vals := sq.Eq{
		"deleted_at": time.Now().Unix(),
	}
	updateBuilder := sq.Update(r.GetTableName()).SetMap(vals)
	return updateBuilder
}
