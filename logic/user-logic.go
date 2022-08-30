package logic

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/response"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	errorhandling "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/error-handling"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type UserLogic interface {
	CreateUser(userRequestObject request.UserRequest) string
	GetUsers() (bool, string, *[]response.UserResponse)
	GetUser(userId int) (bool, string, *response.UserResponse)
	UpdateUser(userRequestObject request.UserRequest, userID int) string
	DeleteUser(userId int) string
}

type UserLogicImpl struct {
	repo repo.Repo
}

func NewUserLogic(repository repo.Repo) UserLogic {
	return &UserLogicImpl{repo: repository}
}

func (userLogic UserLogicImpl) CreateUser(userRequest request.UserRequest) string	{
	user := mapUserRequestToUserDao(userRequest)
	err := userLogic.repo.CreateUser(user)
	if err != nil	{
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}

	log.Println("User created successfully")
	return "User created successfully"
}

func (userLogic UserLogicImpl) GetUsers() (bool, string, *[]response.UserResponse)	{
	var users *[]entities.User
	var err error
	if users, err = userLogic.repo.GetUsers(); err != nil	{
		if err == gorm.ErrRecordNotFound	{
			return false, "No records found", nil
		}
		return false, errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error()), nil
	}
	var usersResponses []response.UserResponse
	for _, user := range *users	{
		userResponse := mapUserDaoToUserResponse(user)
		usersResponses = append(usersResponses, *userResponse)
	}

	log.Println("Users retrieved successfully")
	return true, "", &usersResponses
}

func (userLogic UserLogicImpl) GetUser(userId int) (bool, string, *response.UserResponse)	{
	var user *entities.User
	var err error
	if user, err = userLogic.repo.GetUser(userId); err != nil	{
		if err == gorm.ErrRecordNotFound	{
			return false, "No records found", nil
		}
		return false, errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error()), nil
	}
	userResponse := mapUserDaoToUserResponse(*user)
	log.Println("User", userId, "retrieved successfully")
	return true, "", userResponse
}

func (userLogic UserLogicImpl) UpdateUser(userRequestObject request.UserRequest, userID int) string	{
	user := mapUserRequestToUserDao(userRequestObject)
	if err := userLogic.repo.UpdateUser(user, userID); err != nil	{
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	log.Println("User updated successfully")
	return "User updated successfully"
}

func (userLogic UserLogicImpl) DeleteUser(userId int) string	{
	err := userLogic.repo.DeleteUser(userId)
	if err != nil	{
		return err.Error()
	}
	return "User deleted successfully"
}

func mapUserRequestToUserDao(userRequestObject request.UserRequest) *entities.User {
	user := &entities.User{
		FirstName: userRequestObject.FirstName,
		Email:     userRequestObject.Email,
		LastName:  userRequestObject.LastName,
	}
	return user
}

func mapUserDaoToUserResponse(user entities.User) *response.UserResponse {
	userResponse := &response.UserResponse{
		FirstName: user.FirstName,
		Email:     user.Email,
		LastName:  user.LastName,
	}
	return userResponse
}