package sensor

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	db_models "github.com/krobus00/iot-be/model/database"
)

func Test_repository_Store(t *testing.T) {
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
			name: "WHEN store new sensor data THEN system store sensor data into DB",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    70,
					Temperature: 10,
					HeatIndex:   30,
					DateColumn:  db_models.DateColumn{},
				},
			},
			mock: &mock{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN store new sensor data THEN system return error",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    70,
					Temperature: 10,
					HeatIndex:   30,
					DateColumn:  db_models.DateColumn{},
				},
			},
			mock: &mock{
				err: errors.New("error"),
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
				mockQuery := mock.ExpectExec(`^INSERT INTO sensors \(heat_index,humidity,id,node_id,temperature\) VALUES \(\?,\?,\?,\?,\?\)$`)

				if tt.mock.err == nil {
					mockQuery.WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			if err := r.Store(tt.args.ctx, sqlxDB, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("repository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
