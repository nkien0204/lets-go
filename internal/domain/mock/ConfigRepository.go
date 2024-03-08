// Code generated by mockery v2.35.4. DO NOT EDIT.

package mock

import (
	config "github.com/nkien0204/lets-go/internal/domain/entity/config"

	mock "github.com/stretchr/testify/mock"
)

// ConfigRepository is an autogenerated mock type for the ConfigRepository type
type ConfigRepository struct {
	mock.Mock
}

// ReadConfigFile provides a mock function with given fields:
func (_m *ConfigRepository) ReadConfigFile() (config.ConfigFileReadResponseEntity, error) {
	ret := _m.Called()

	var r0 config.ConfigFileReadResponseEntity
	var r1 error
	if rf, ok := ret.Get(0).(func() (config.ConfigFileReadResponseEntity, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() config.ConfigFileReadResponseEntity); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.ConfigFileReadResponseEntity)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewConfigRepository creates a new instance of ConfigRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigRepository {
	mock := &ConfigRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
