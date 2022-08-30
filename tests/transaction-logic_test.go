package tests

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var transactionMockRepo = new(repo.MockRepository)
var transactionMockCache = new(cache.MockCache)
var transactionLogic = logic.NewTransactionLogic(transactionMockRepo, transactionMockCache)

func TestRequestTransaction(t *testing.T) {
	transactionRequest := &request.TransactionRequest{
		WalletId:        1,
		Amount:          100,
		TransactionType: "CREDIT",
	}
	wallet := &entities.Wallet{
		UserId:  6,
		Balance: 100.98,
		Status:  "ACTIVE",
	}
	transactionMockRepo.On("FindWalletForUpdate").Return(wallet, nil)
	transactionMockRepo.On("SaveWallet").Return(nil)
	transactionMockRepo.On("SaveTransaction").Return(nil)
	transactionMockCache.On("SET").Return(nil)
	result := transactionLogic.RequestTransaction(transactionRequest)

	transactionMockRepo.AssertExpectations(t)
	transactionMockCache.AssertExpectations(t)
	assert.Equal(t, "Transaction Successful", result)
}