package sensor

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
)

func Test_repository_GetSensorByRange(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *model.GetProcessedDataRequest
	}

	type mock struct {
		res []*db_models.Sensor
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    *mock
		want    []*db_models.Sensor
		wantErr bool
	}{
		{
			name: "WHEN get sensor data between range THEN system return sensors data",
			args: args{
				ctx: context.TODO(),
				input: &model.GetProcessedDataRequest{
					NodeID:    "NODE-UUID",
					StartDate: 1653753862,
					EndDate:   1653753862,
				},
			},
			mock: &mock{
				res: []*db_models.Sensor{
					{
						ID:          "UUID",
						NodeID:      "NODE-UUID",
						Humidity:    70,
						Temperature: 10,
						HeatIndex:   30,
						DateColumn:  db_models.DateColumn{},
					},
					{
						ID:          "UUID2",
						NodeID:      "NODE-UUID",
						Humidity:    70,
						Temperature: 10,
						HeatIndex:   30,
						DateColumn:  db_models.DateColumn{},
					},
				},
				err: nil,
			},
			want: []*db_models.Sensor{
				{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    70,
					Temperature: 10,
					HeatIndex:   30,
					DateColumn:  db_models.DateColumn{},
				},
				{
					ID:          "UUID2",
					NodeID:      "NODE-UUID",
					Humidity:    70,
					Temperature: 10,
					HeatIndex:   30,
					DateColumn:  db_models.DateColumn{},
				},
			},
			wantErr: false,
		},
		{
			name: "WHEN get sensor data between range AND data not found THEN system empty array",
			args: args{
				ctx: context.TODO(),
				input: &model.GetProcessedDataRequest{
					NodeID:    "NODE-UUID",
					StartDate: 1653753862,
					EndDate:   1653753862,
				},
			},
			mock: &mock{
				res: nil,
				err: sql.ErrNoRows,
			},
			want:    []*db_models.Sensor{},
			wantErr: false,
		},
		{
			name: "WHEN get sensor data between range THEN system return error",
			args: args{
				ctx: context.TODO(),
				input: &model.GetProcessedDataRequest{
					NodeID:    "NODE-UUID",
					StartDate: 1653753862,
					EndDate:   1653753862,
				},
			},
			mock: &mock{
				res: nil,
				err: errors.New("error"),
			},
			want:    nil,
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
				mockQuery := mock.ExpectQuery(`^SELECT id, node_id, humidity, temperature, heat_index, created_at, updated_at, deleted_at FROM sensors WHERE deleted_at IS NULL AND \(node_id = \? AND created_at >= \? AND created_at <= \?\) ORDER BY created_at ASC$`)

				if tt.mock.err == nil {
					rows := sqlmock.NewRows([]string{"id", "node_id", "humidity", "temperature", "heat_index", "created_at", "updated_at", "deleted_at"})
					for _, row := range tt.mock.res {
						rows.AddRow(row.ID, row.NodeID, row.Humidity, row.Temperature, row.HeatIndex, row.CreatedAt, row.UpdatedAt, row.DeletedAt)
					}
					mockQuery.WillReturnRows(rows)
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			got, err := r.GetSensorByRange(tt.args.ctx, sqlxDB, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetSensorByRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetSensorByRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
