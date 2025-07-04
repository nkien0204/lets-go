// Code generated by mockery v2.20.0. DO NOT EDIT.

package mock

import (
	domain "github.com/nkien0204/lets-go/internal/domain"
	generator "github.com/nkien0204/lets-go/internal/domain/entity/generator"
	mock "github.com/stretchr/testify/mock"
)

// GeneratorDelivery is an autogenerated mock type for the GeneratorDelivery type
type GeneratorDelivery struct {
	mock.Mock
}

// HandleOffGenerate provides a mock function with given fields: _a0
func (_m *GeneratorDelivery) HandleOffGenerate(_a0 generator.GeneratorInputEntity) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(generator.GeneratorInputEntity) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleOnlGenerate provides a mock function with given fields: _a0
func (_m *GeneratorDelivery) HandleOnlGenerate(_a0 generator.GeneratorInputEntity) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(generator.GeneratorInputEntity) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOfflineUsecase provides a mock function with given fields: _a0
func (_m *GeneratorDelivery) SetOfflineUsecase(_a0 domain.GeneratorUsecase) {
	_m.Called(_a0)
}

// SetOnlineUsecase provides a mock function with given fields: _a0
func (_m *GeneratorDelivery) SetOnlineUsecase(_a0 domain.GeneratorUsecase) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewGeneratorDelivery interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneratorDelivery creates a new instance of GeneratorDelivery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneratorDelivery(t mockConstructorTestingTNewGeneratorDelivery) *GeneratorDelivery {
	mock := &GeneratorDelivery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
