package server

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitializeV1Routers()	{
	r :=  mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	// swagger:operation POST /accounts/ accounts createAccount
	// ---
	// summary: Creates a new account.
	// description: If account creation is success, account will be returned with Created (201).
		// parameters:
		// - name: account
		//   description: account to add to the list of accounts
		//   in: body
		//   required: true
		//   schema:
		//     "$ref": entities.User
		// responses:
		//   "200":
		//     "$ref": "User Created Successfully"
		//   "400":
		//     "$ref": "Something went wrong"
	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users", api.GetUsers).Methods("GET")
	r.HandleFunc("/users/{user_id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/{user_id}", api.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{user_id}", api.DeleteUser).Methods("DELETE")

	r.HandleFunc("/users/wallets", api.CreateWallet).Methods("POST")
	r.HandleFunc("/users/wallets/{wallet_id}/balance", api.GetBalance).Methods("GET")
	r.HandleFunc("/users/wallets/status", api.RequestWalletStatusChange).Methods("PUT")
	r.HandleFunc("/users/wallets/transaction", api.RequestTransaction).Methods("POST")
	r.HandleFunc("/users/wallets/{wallet_id}", api.DeleteWallet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}