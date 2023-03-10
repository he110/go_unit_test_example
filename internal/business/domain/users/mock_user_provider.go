// Code generated by mockery. DO NOT EDIT.

package users

import mock "github.com/stretchr/testify/mock"

// mockUserProvider is an autogenerated mock type for the userProvider type
type mockUserProvider struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *mockUserProvider) FindAll() ([]User, error) {
	ret := _m.Called()

	var r0 []User
	if rf, ok := ret.Get(0).(func() []User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockUserProvider interface {
	mock.TestingT
	Cleanup(func())
}

// newMockUserProvider creates a new instance of mockUserProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockUserProvider(t mockConstructorTestingTnewMockUserProvider) *mockUserProvider {
	mock := &mockUserProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
