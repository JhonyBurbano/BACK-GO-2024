package service

import (
	"context"
	"errors"
	"github.com/jnates/smartOshApi/domain/users/domain/model"
	"github.com/jnates/smartOshApi/mocks"
	response "github.com/jnates/smartOshApi/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	assertions := assert.New(t)
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo)

	user := &model.User{
		PersonaID: 1,
	}

	expectedResponse := &response.CreateResponse{
		Message: "User created",
	}

	mockRepo.On("CreateUser", context.Background(), user).Return(expectedResponse, nil)

	res, err := userService.CreateUser(context.Background(), user)
	assertions.NoError(err)
	assertions.Equal(expectedResponse, res)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "CreateUser", 1)
	mockRepo.AssertCalled(t, "CreateUser", context.Background(), user)
}

func TestGetUserNotFound(t *testing.T) {
	assertions := assert.New(t)
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo)

	expectedResponse := &response.GenericUserResponse{
		Error: "User not found",
	}

	mockRepo.On("GetUser", context.Background(), "2").Return(expectedResponse, errors.New("user not found"))

	res, err := userService.GetUser(context.Background(), "2")
	assertions.Error(err)
	assertions.Equal(expectedResponse, res)
}

func TestGetUser(t *testing.T) {
	assertions := assert.New(t)
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo)

	user := &model.User{
		PersonaID: 1,
	}

	expectedResponse := &response.GenericUserResponse{
		Message: "Get user success",
		User:    user,
	}

	mockRepo.On("GetUser", context.Background(), "1").Return(expectedResponse, nil)

	res, err := userService.GetUser(context.Background(), "1")
	assertions.NoError(err)
	assertions.Equal(expectedResponse, res)
}

func TestGetUsers(t *testing.T) {
	assertions := assert.New(t)
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo)

	user1 := &model.User{
		PersonaID: 1,
	}

	user2 := &model.User{
		PersonaID: 2,
	}

	expectedResponse := &response.GenericUserResponse{
		Message: "Get users success",
		User:    []*model.User{user1, user2},
	}

	mockRepo.On("GetUsers", context.Background()).Return(expectedResponse, nil)

	res, err := userService.GetUsers(context.Background())
	assertions.NoError(err)
	assertions.Equal(expectedResponse, res)
}

func TestLoginUser(t *testing.T) {
	assertions := assert.New(t)
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo)

	user := &model.User{
		PersonaID: 1,
	}

	expectedResponse := &response.GenericUserResponse{
		Message: "Login success",
		User:    user,
	}

	mockRepo.On("LoginUser", context.Background(), user).Return(expectedResponse, nil)

	res, err := userService.LoginUser(context.Background(), user)
	assertions.NoError(err)
	assertions.Equal(expectedResponse, res)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "LoginUser", 1)
	mockRepo.AssertCalled(t, "LoginUser", context.Background(), user)
}
