package service

import (
	"app/internal"
	"github.com/stretchr/testify/mock"
)

func NewMocksVehicle() *MockVehicle {
	return &MockVehicle{}
}

type MockVehicle struct {
	mock.Mock
}

func (m *MockVehicle) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *MockVehicle) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *MockVehicle) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	args := m.Called(brand)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockVehicle) AverageCapacityByBrand(brand string) (a int, err error) {
	args := m.Called(brand)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockVehicle) SearchByWeightRange(query internal.SearchQuery, b bool) (v map[int]internal.Vehicle, err error) {
	args := m.Called(query, b)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
