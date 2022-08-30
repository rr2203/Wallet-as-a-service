package api

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	errorhandling "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/error-handling"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"encoding/json"
	"net/http"
)

var transactionLogic = logic.NewTransactionLogic(repo.NewMySQLRepo(), cache.NewRedisCache("localhost:6379", 0, 0))

func RequestTransaction(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	var transactionRequest request.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		userErrorMessage := errorhandling.HandleError(http.StatusBadRequest, "Invalid Input", err.Error())
		json.NewEncoder(w).Encode(userErrorMessage)
		return
	}
	transactionRequestResult := transactionLogic.RequestTransaction(&transactionRequest)
	json.NewEncoder(w).Encode(transactionRequestResult)
}