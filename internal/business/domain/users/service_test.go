package users

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	description string
	service     *Service
	login       string
	assert      func(*testing.T, bool, error)
}

func testIsUserExistsFound() testCase {
	var (
		description    = "IsUserExists User found"
		requestedLogin = "helpdeskboy"
	)

	users := []User{
		{Login: "helpdeskboy"},
		{Login: "ehot_riss"},
	}

	providerMock := new(mockUserProvider)
	providerMock.On("FindAll").Once().Return(users, nil)

	return testCase{
		description: description,
		service:     NewService(providerMock),
		login:       requestedLogin,
		assert: func(t *testing.T, result bool, err error) {
			assert.NoError(t, err)
			assert.True(t, result)
		},
	}
}

func testIsUserExistsNotFound() testCase {
	var (
		description    = "IsUserExists User not found"
		requestedLogin = "play_code"
	)

	users := []User{
		{Login: "helpdeskboy"},
		{Login: "ehot_riss"},
	}

	providerMock := new(mockUserProvider)
	providerMock.On("FindAll").Once().Return(users, nil)

	return testCase{
		description: description,
		service:     NewService(providerMock),
		login:       requestedLogin,
		assert: func(t *testing.T, result bool, err error) {
			assert.NoError(t, err)
			assert.False(t, result)
		},
	}
}

func testIsUserExistsProviderError() testCase {
	var (
		description    = "IsUserExists Provider error"
		requestedLogin = "helpdeskboy"
	)

	providerMock := new(mockUserProvider)
	providerMock.On("FindAll").Once().Return([]User{}, errors.New("some error"))

	return testCase{
		description: description,
		service:     NewService(providerMock),
		login:       requestedLogin,
		assert: func(t *testing.T, result bool, err error) {
			assert.Error(t, err)
			assert.False(t, result)
		},
	}
}

func TestService_IsUserExists(t *testing.T) {
	testCases := []testCase{
		testIsUserExistsFound(),
		testIsUserExistsNotFound(),
		testIsUserExistsProviderError(),
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := tc.service.IsUserExists(tc.login)
			tc.assert(t, result, err)
		})
	}
}
