package tests

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userMockRepo = new(repo.MockRepository)
var userLogic = logic.NewUserLogic(userMockRepo)

func TestCreateUser(t *testing.T) {
	userRequest := request.UserRequest{
		FirstName: "Rahul",
		LastName: "Raj",
		Email: "a@b.com",
	}
	userMockRepo.On("CreateUser").Return(nil)
	result := userLogic.CreateUser(userRequest)

	userMockRepo.AssertExpectations(t)
	assert.Equal(t, "User created successfully", result)
}

func TestGetUsers(t *testing.T) {
	user1 := entities.User{
		FirstName: "Rahul",
		LastName: "Raj",
		Email: "a@b.com",
	}
	user2 := entities.User{
		FirstName: "Rohit",
		LastName: "Raj",
		Email: "c@b.com",
	}
	userMockRepo.On("GetUsers").Return(&[]entities.User{user1, user2}, nil)
	success, message, usersResponses := userLogic.GetUsers()

	userMockRepo.AssertExpectations(t)
	assert.Equal(t, "Rahul", (*usersResponses)[0].FirstName)
	assert.Equal(t, "Raj", (*usersResponses)[0].LastName)
	assert.Equal(t, "a@b.com", (*usersResponses)[0].Email)
	assert.Equal(t, "Rohit", (*usersResponses)[1].FirstName)
	assert.Equal(t, "Raj", (*usersResponses)[1].LastName)
	assert.Equal(t, "c@b.com", (*usersResponses)[1].Email)
	assert.Equal(t, true, success)
	assert.Equal(t, message, "Users retrieved successfully")
}

func TestGetUser(t *testing.T) {
	user := &entities.User{
		FirstName: "Rahul",
		LastName: "Raj",
		Email: "a@b.com",
	}
	userMockRepo.On("GetUser").Return(user, nil)
	success, message, userResponse := userLogic.GetUser(1)

	userMockRepo.AssertExpectations(t)
	assert.Equal(t, true, success)
	assert.Equal(t, message, "User retrieved successfully")
	assert.Equal(t, "Rahul", userResponse.FirstName)
	assert.Equal(t, "Raj", userResponse.LastName)
	assert.Equal(t, "a@b.com", userResponse.Email)
}

func TestUpdateUser(t *testing.T) {
	userRequest := request.UserRequest{
		FirstName: "Rahul",
		LastName: "Raj",
		Email: "a@b.com",
	}
	userMockRepo.On("UpdateUser").Return(nil)
	result := userLogic.UpdateUser(userRequest, 1)

	userMockRepo.AssertExpectations(t)
	assert.Equal(t, "User updated successfully", result)
}

func TestDeleteUser(t *testing.T) {
	userMockRepo.On("DeleteUser").Return(nil)
	result := userLogic.DeleteUser(1)

	userMockRepo.AssertExpectations(t)
	assert.Equal(t, "User deleted successfully", result)
}