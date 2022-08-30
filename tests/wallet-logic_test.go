package tests

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var walletMockRepo = new(repo.MockRepository)
var walletMockCache = new(cache.MockCache)
var walletLogic = logic.NewWalletLogic(walletMockRepo, walletMockCache)

func TestCreateWallet(t *testing.T) {
	walletRequest := request.WalletRequest{
		UserId: 1,
		Balance: 100.98,
		Status: "ACTIVE",
	}
	walletMockRepo.On("CreateWallet").Return(nil)
	result := walletLogic.CreateWallet(walletRequest)

	walletMockRepo.AssertExpectations(t)
	assert.Equal(t, "Wallet created successfully", result)
}

func TestGetBalance(t *testing.T) {
	wallet := &entities.Wallet{
		UserId:  6,
		Balance: 100.98,
		Status:  "ACTIVE",
	}
	walletMockRepo.On("FindWallet").Return(wallet, nil)
	walletMockCache.On("GET").Return("", errors.New("key not found"))
	walletMockCache.On("SET").Return(nil)
	balance, err := walletLogic.GetBalance(6)

	walletMockRepo.AssertExpectations(t)
	walletMockCache.AssertExpectations(t)
	assert.Equal(t, nil, err)
	assert.Equal(t, 100.98, balance)
}

func TestRequestWalletStatusChange(t *testing.T) {
	wallet := &entities.Wallet{
		UserId:  6,
		Balance: 100.98,
		Status:  "ACTIVE",
	}
	walletMockRepo.On("FindWallet").Return(wallet, nil)
	walletMockRepo.On("SaveWallet").Return(nil)
	result := walletLogic.RequestWalletStatusChange(1, "ACTIVE")

	walletMockRepo.AssertExpectations(t)
	assert.Equal(t, "Wallet status changed successfully", result)
}

func TestDeleteWallet(t *testing.T) {
	walletMockRepo.On("DeleteWallet").Return(nil)
	result := walletLogic.DeleteWallet(1)

	walletMockRepo.AssertExpectations(t)
	assert.Equal(t, "Wallet deleted successfully", result)
}