package api

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	errorhandling "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/error-handling"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var userLogic = logic.NewUserLogic(repo.NewMySQLRepo())

func CreateUser(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	var userRequest request.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest, "Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	result := userLogic.CreateUser(userRequest)
	json.NewEncoder(w).Encode(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	success, message, users := userLogic.GetUsers()
	if success {
		json.NewEncoder(w).Encode(users)
	} else	{
		json.NewEncoder(w).Encode(message)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	success, message, user := userLogic.GetUser(userID)

	if success {
		json.NewEncoder(w).Encode(user)
	} else	{
		json.NewEncoder(w).Encode(message)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	var user request.UserRequest
	json.NewDecoder(r.Body).Decode(&user)
	result := userLogic.UpdateUser(user, userID)
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	result := userLogic.DeleteUser(userID)
	json.NewEncoder(w).Encode(result)
}
