package data

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Search(name string) []Kitten {
	args := m.Mock.Called(name)

	return args.Get(0).([]Kitten)
}
