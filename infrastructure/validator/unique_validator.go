package validator

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator/v10"
)

func (cv *customValidator) UniqueValidator(fl validator.FieldLevel) bool {

	ctx := context.Background()
	var count int64
	params := strings.Split(fl.Param(), " ")

	if len(params) < 2 || len(params) > 3 {
		return false
	}
	tableName := params[0]
	tableColumn := params[1]

	countBuilder := sq.Select("count(*)").Where(sq.Eq{tableColumn: fl.Field().String()}).From(tableName)
	if len(params) == 3 && params[2] == "exclude" {
		countBuilder = countBuilder.Where(sq.NotEq{tableColumn: fl.Field().String()})
	}
	countQuery, args, err := countBuilder.ToSql()
	if err != nil {
		return false
	}
	err = cv.db.SqlxDB.GetContext(ctx, &count, countQuery, args...)
	if err != nil {

		return false
	}

	return count == 0
}
