// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	domains "week5/internal/core/domains"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockStudentRepo is an autogenerated mock type for the StudentRepo type
type MockStudentRepo struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ID
func (_m *MockStudentRepo) GetByID(ID uint) (*domains.Student, error) {
	ret := _m.Called(ID)

	var r0 *domains.Student
	if rf, ok := ret.Get(0).(func(uint) *domains.Student); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: student
func (_m *MockStudentRepo) Save(student *domains.Student) error {
	ret := _m.Called(student)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domains.Student) error); ok {
		r0 = rf(student)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockStudentRepo creates a new instance of MockStudentRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStudentRepo(t testing.TB) *MockStudentRepo {
	mock := &MockStudentRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
