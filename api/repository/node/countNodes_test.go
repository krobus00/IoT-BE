package node

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
)

func Test_repository_CountNodes(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type mock struct {
		res int64
		err error
	}
	tests := []struct {
		name    string
		args    args
		mock    *mock
		want    int64
		wantErr bool
	}{
		{
			name: "WHEN count rows THEN system return total row",
			args: args{
				ctx: context.TODO(),
			},
			mock: &mock{
				res: 17,
				err: nil,
			},
			want:    17,
			wantErr: false,
		},
		{
			name: "WHEN count rows THEN system return error",
			args: args{
				ctx: context.TODO(),
			},
			mock: &mock{
				res: 0,
				err: errors.New("error"),
			},
			want:    0,
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
				mockQuery := mock.ExpectQuery(`^SELECT count\(id\) FROM nodes WHERE deleted_at IS NULL$`)

				if tt.mock.err == nil {
					rows := sqlmock.NewRows([]string{"count"})
					row := tt.mock.res
					rows.AddRow(row)
					mockQuery.WillReturnRows(rows)
				}

				mockQuery.WillReturnError(tt.mock.err)
			}
			got, err := r.CountNodes(tt.args.ctx, sqlxDB)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CountNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.CountNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
