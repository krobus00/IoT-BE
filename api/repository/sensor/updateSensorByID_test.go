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

func Test_repository_UpdateSensorByID(t *testing.T) {

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
			name: "WHEN update sensor data by ID THEN data updated",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   10,
					DateColumn:  db_models.DateColumn{},
				},
			},
			mock: &mock{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN update sensor data by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   10,
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
				mockQuery := mock.ExpectExec(`^UPDATE sensors SET heat_index = \?, humidity = \?, node_id = \?, temperature = \?, updated_at = \? WHERE id = \?$`)

				if tt.mock.err == nil {
					mockQuery.WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			if err := r.UpdateSensorByID(tt.args.ctx, sqlxDB, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateSensorByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
