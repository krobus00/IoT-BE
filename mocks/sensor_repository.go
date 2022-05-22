// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	iot_bemodel "github.com/krobus00/iot-be/model"
	database "github.com/krobus00/iot-be/model/database"

	mock "github.com/stretchr/testify/mock"

	model "github.com/krobus00/krobot-building-block/model"

	sqlx "github.com/jmoiron/sqlx"

	testing "testing"
)

// SensorRepository is an autogenerated mock type for the SensorRepository type
type SensorRepository struct {
	mock.Mock
}

// DeleteSensorByID provides a mock function with given fields: ctx, db, input
func (_m *SensorRepository) DeleteSensorByID(ctx context.Context, db *sqlx.DB, input *database.Sensor) error {
	ret := _m.Called(ctx, db, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *database.Sensor) error); ok {
		r0 = rf(ctx, db, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllSensor provides a mock function with given fields: ctx, db, paginationRequest, config
func (_m *SensorRepository) GetAllSensor(ctx context.Context, db *sqlx.DB, paginationRequest *model.PaginationRequest, config ...model.Config) ([]*database.Sensor, int64, error) {
	_va := make([]interface{}, len(config))
	for _i := range config {
		_va[_i] = config[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, db, paginationRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*database.Sensor
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *model.PaginationRequest, ...model.Config) []*database.Sensor); ok {
		r0 = rf(ctx, db, paginationRequest, config...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.Sensor)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, *model.PaginationRequest, ...model.Config) int64); ok {
		r1 = rf(ctx, db, paginationRequest, config...)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *sqlx.DB, *model.PaginationRequest, ...model.Config) error); ok {
		r2 = rf(ctx, db, paginationRequest, config...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLastReportByNodeID provides a mock function with given fields: ctx, db, input
func (_m *SensorRepository) GetLastReportByNodeID(ctx context.Context, db *sqlx.DB, input *database.Sensor) (*database.Sensor, error) {
	ret := _m.Called(ctx, db, input)

	var r0 *database.Sensor
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *database.Sensor) *database.Sensor); ok {
		r0 = rf(ctx, db, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Sensor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, *database.Sensor) error); ok {
		r1 = rf(ctx, db, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSensorByID provides a mock function with given fields: ctx, db, input
func (_m *SensorRepository) GetSensorByID(ctx context.Context, db *sqlx.DB, input *database.Sensor) (*database.Sensor, error) {
	ret := _m.Called(ctx, db, input)

	var r0 *database.Sensor
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *database.Sensor) *database.Sensor); ok {
		r0 = rf(ctx, db, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Sensor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, *database.Sensor) error); ok {
		r1 = rf(ctx, db, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSensorByRange provides a mock function with given fields: ctx, db, input
func (_m *SensorRepository) GetSensorByRange(ctx context.Context, db *sqlx.DB, input *iot_bemodel.GetProcessedDataRequest) ([]*database.Sensor, error) {
	ret := _m.Called(ctx, db, input)

	var r0 []*database.Sensor
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *iot_bemodel.GetProcessedDataRequest) []*database.Sensor); ok {
		r0 = rf(ctx, db, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.Sensor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, *iot_bemodel.GetProcessedDataRequest) error); ok {
		r1 = rf(ctx, db, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTableName provides a mock function with given fields:
func (_m *SensorRepository) GetTableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Store provides a mock function with given fields: ctx, db, input
func (_m *SensorRepository) Store(ctx context.Context, db *sqlx.DB, input *database.Sensor) error {
	ret := _m.Called(ctx, db, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *database.Sensor) error); ok {
		r0 = rf(ctx, db, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSensorByID provides a mock function with given fields: ctx, db, input
func (_m *SensorRepository) UpdateSensorByID(ctx context.Context, db *sqlx.DB, input *database.Sensor) error {
	ret := _m.Called(ctx, db, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *database.Sensor) error); ok {
		r0 = rf(ctx, db, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSensorRepository creates a new instance of SensorRepository. It also registers a cleanup function to assert the mocks expectations.
func NewSensorRepository(t testing.TB) *SensorRepository {
	mock := &SensorRepository{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}