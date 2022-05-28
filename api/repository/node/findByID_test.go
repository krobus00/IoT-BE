package node

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	db_models "github.com/krobus00/iot-be/model/database"
)

func Test_repository_FindNodeByID(t *testing.T) {

	type args struct {
		ctx   context.Context
		input *db_models.Node
	}
	type mock struct {
		res *db_models.Node
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    *mock
		want    *db_models.Node
		wantErr bool
	}{
		{
			name: "WHEN get node by ID THEN system return node data",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10.10,
					Latitude:   17.17,
					ModelURL:   "https://bucket_url",
					DateColumn: db_models.DateColumn{},
				},
			},
			mock: &mock{
				res: &db_models.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10.10,
					Latitude:   17.17,
					ModelURL:   "https://bucket_url",
					DateColumn: db_models.DateColumn{},
				},
				err: nil,
			},
			want: &db_models.Node{
				ID:         "UUID",
				City:       "city",
				Longitude:  10.10,
				Latitude:   17.17,
				ModelURL:   "https://bucket_url",
				DateColumn: db_models.DateColumn{},
			},
			wantErr: false,
		},
		{
			name: "WHEN get node by ID AND data not found THEN system return nil",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10.10,
					Latitude:   17.17,
					ModelURL:   "https://bucket_url",
					DateColumn: db_models.DateColumn{},
				},
			},
			mock: &mock{
				res: nil,
				err: sql.ErrNoRows,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "WHEN get node by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				input: &db_models.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10.10,
					Latitude:   17.17,
					ModelURL:   "https://bucket_url",
					DateColumn: db_models.DateColumn{},
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
				mockQuery := mock.ExpectQuery(`^SELECT id, city, longitude, latitude, created_at, updated_at, deleted_at, COALESCE\(nodes.model_url,nodes.fallback_model_url\) AS model_url FROM nodes WHERE deleted_at IS NULL AND id = \?$`)

				if tt.mock.err == nil {
					rows := sqlmock.NewRows([]string{"id", "city", "longitude", "latitude", "model_url", "created_at", "updated_at", "deleted_at"})
					row := tt.mock.res
					rows.AddRow(row.ID, row.City, row.Longitude, row.Latitude, row.ModelURL, row.CreatedAt, row.UpdatedAt, row.DeletedAt)
					mockQuery.WillReturnRows(rows)
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			got, err := r.FindNodeByID(tt.args.ctx, sqlxDB, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.FindNodeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.FindNodeByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
