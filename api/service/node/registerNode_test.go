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
	"github.com/stretchr/testify/mock"
)

func Test_service_Register(t *testing.T) {

	type args struct {
		ctx     context.Context
		payload *model.RegisterRequest
	}
	type mockStore struct {
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockStore *mockStore
		wantErr   bool
	}{
		{
			name: "WHEN regisner new node THEN system store new node into DB and return node ID",
			args: args{
				ctx: context.TODO(),
				payload: &model.RegisterRequest{
					City:      "city",
					Longitude: 10,
					Latitude:  17,
				},
			},
			mockStore: &mockStore{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN regisner new node THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.RegisterRequest{
					City:      "city",
					Longitude: 10,
					Latitude:  17,
				},
			},
			mockStore: &mockStore{
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

			if tt.mockStore != nil {
				mockNodeRepo.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockStore.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					NodeRepository: mockNodeRepo,
				},
			}
			_, err := svc.Register(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
