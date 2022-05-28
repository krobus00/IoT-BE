package sensor

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	db_models "github.com/krobus00/iot-be/model/database"
)

func Test_repository_DeleteSensorByID(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *db_models.Sensor
	}
	type mock struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    *mock
		wantErr bool
	}{
		{
			name: "WHEN delete sensor data by ID THEN deleted_at updated",
			args: args{
				ctx:   context.TODO(),
				input: &db_models.Sensor{ID: "UUID"},
			},
			mock: &mock{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN delete sensor data by ID AND data not found THEN system return error",
			args: args{
				ctx:   context.TODO(),
				input: &db_models.Sensor{ID: "UUID"},
			},
			mock: &mock{
				err: sql.ErrNoRows,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")
			defer sqlxDB.Close()

			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)
			r := &repository{
				logger: logger,
			}
			if tt.mock != nil {
				mockQuery := mock.ExpectExec(`^UPDATE sensors SET deleted_at = \? WHERE id = \?$`)

				if tt.mock.err == nil {
					mockQuery.WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			if err := r.DeleteSensorByID(tt.args.ctx, sqlxDB, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteSensorByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
