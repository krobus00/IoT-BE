package sensor

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

func Test_repository_GetAllSensor(t *testing.T) {
	type args struct {
		ctx               context.Context
		paginationRequest *kro_model.PaginationRequest
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
			name: "WHEN get all sensor THEN system return paginated data",
			args: args{
				ctx:               context.TODO(),
				paginationRequest: &kro_model.PaginationRequest{},
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
			name: "WHEN get all sensor THEN system return error",
			args: args{
				ctx:               context.TODO(),
				paginationRequest: &kro_model.PaginationRequest{},
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
				mockQuery := mock.ExpectQuery(`^SELECT id, node_id, humidity, temperature, heat_index, created_at, updated_at, deleted_at FROM sensors WHERE deleted_at IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 0$`)

				if tt.mock.err == nil {
					rows := sqlmock.NewRows([]string{"id", "node_id", "humidity", "temperature", "heat_index", "created_at", "updated_at", "deleted_at"})
					for _, row := range tt.mock.res {
						rows.AddRow(row.ID, row.NodeID, row.Humidity, row.Temperature, row.HeatIndex, row.CreatedAt, row.UpdatedAt, row.DeletedAt)
					}
					mockQuery.WillReturnRows(rows)
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			got, err := r.GetAllSensor(tt.args.ctx, sqlxDB, tt.args.paginationRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllSensor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAllSensor() = %v, want %v", got, tt.want)
			}
		})
	}
}
