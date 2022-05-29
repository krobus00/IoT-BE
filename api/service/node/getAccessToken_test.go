package node

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	"github.com/stretchr/testify/mock"
)

func Test_service_GetAccessToken(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *model.GetAccessTokenRequest
	}
	type mockFindNodeByID struct {
		res *database.Node
		err error
	}
	tests := []struct {
		name             string
		args             args
		mockFindNodeByID *mockFindNodeByID
		wantErr          bool
	}{
		{
			name: "WHEN get access token THEN system return node access token",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetAccessTokenRequest{
					ID: "UUID",
				},
			},
			mockFindNodeByID: &mockFindNodeByID{
				res: &database.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10.17,
					Latitude:   17.10,
					ModelURL:   "https://bucket-name",
					DateColumn: database.DateColumn{},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN get access token AND node not found THEN system error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetAccessTokenRequest{
					ID: "UUID",
				},
			},
			mockFindNodeByID: &mockFindNodeByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "WHEN get access token THEN system error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetAccessTokenRequest{
					ID: "UUID",
				},
			},
			mockFindNodeByID: &mockFindNodeByID{
				res: nil,
				err: errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")
			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)

			mockNodeRepo := new(mocks.NodeRepository)

			if tt.mockFindNodeByID != nil {
				mockNodeRepo.On("FindNodeByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFindNodeByID.res, tt.mockFindNodeByID.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					NodeRepository: mockNodeRepo,
				},
			}
			_, err := svc.GetAccessToken(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
