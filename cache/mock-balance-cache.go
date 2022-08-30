package cache

import (
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (mock *MockCache) SET(int, float64) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockCache) GET(int) (string, error) {
	args := mock.Called()
	return args.Get(0).(string), args.Error(1)
}