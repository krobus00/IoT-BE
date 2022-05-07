package node

import (
	sq "github.com/Masterminds/squirrel"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

var (
	searchFields = []string{"city", "longitude", "latitude"}
	whiteList    = kro_model.ColumnWhiteList{
		"id":         true,
		"city":       true,
		"longitude":  true,
		"latitude":   true,
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
	}
	columnMapping = kro_model.ColumnMapping{
		"id":        "id",
		"city":      "city",
		"longitude": "longitude",
		"latitude":  "latitude",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
		"deletedAt": "deleted_at",
	}
)

func (r *repository) buildInsertQuery(input *db_models.Node) sq.InsertBuilder {
	vals := sq.Eq{
		"id":        input.ID,
		"city":      input.City,
		"longitude": input.Longitude,
		"latitude":  input.Latitude,
	}
	insertBuilder := sq.Insert(r.GetTableName()).SetMap(vals)
	return insertBuilder
}

func (r *repository) buildSelectQuery() sq.SelectBuilder {
	selection := []string{
		"id",
		"city",
		"longitude",
		"latitude",
		"created_at",
		"updated_at",
		"deleted_at",
	}
	selectBuilder := sq.Select(selection...).Where(sq.Eq{"deleted_at": nil}).From(r.GetTableName())
	return selectBuilder
}
