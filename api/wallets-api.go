package api

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	errorhandling "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/error-handling"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var walletLogic = logic.NewWalletLogic(repo.NewMySQLRepo(), cache.NewRedisCache("localhost:6379", 0, 0))

func CreateWallet(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	var wallet request.WalletRequest
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	result := walletLogic.CreateWallet(wallet)
	json.NewEncoder(w).Encode(result)
}

func RequestWalletStatusChange(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	var statusChangRequest request.StatusChangeRequest
	err := json.NewDecoder(r.Body).Decode(&statusChangRequest)
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	result := walletLogic.RequestWalletStatusChange(statusChangRequest.WalletID, statusChangRequest.Status)
	json.NewEncoder(w).Encode(result)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	walletID, err := strconv.Atoi(params["wallet_id"])
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	result := walletLogic.DeleteWallet(walletID)
	json.NewEncoder(w).Encode(result)
}

func GetBalance(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	walletID, err := strconv.Atoi(params["wallet_id"])
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest,"Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	var balance float64
	balance, err = walletLogic.GetBalance(walletID)
	if err != nil	{
		userErrorMessage := errorhandling.HandleError(http.StatusInternalServerError,"Failed to retrieve balance", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
	}	else {
		json.NewEncoder(w).Encode(balance)
	}
}