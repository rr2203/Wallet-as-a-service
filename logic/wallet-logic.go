package logic

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	walletCommons "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/commons/wallet"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	errorhandling "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/error-handling"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type WalletLogic interface {
	CreateWallet(walletRequestObject request.WalletRequest) string
	GetBalance(walletID int) (float64, error)
	RequestWalletStatusChange(walletID int, status string)	string
	DeleteWallet(walletID int) string
}

type WalletLogicImpl struct {
	repo repo.Repo
	balanceCache cache.BalanceCache
}

func NewWalletLogic(repo repo.Repo, balanceCache cache.BalanceCache) WalletLogic {
	return &WalletLogicImpl{
		repo:  repo,
		balanceCache: balanceCache,
	}
}

func (walletLogic WalletLogicImpl) CreateWallet(walletRequestObject request.WalletRequest) string	{
	wallet := mapWalletRequestToWalletDao(walletRequestObject)
	err := walletLogic.repo.CreateWallet(wallet)
	if err != nil {
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	log.Println("Wallet created successfully")
	return "Wallet created successfully"
}

func (walletLogic WalletLogicImpl) GetBalance(walletID int) (float64, error) {
	cachedValue, err := walletLogic.balanceCache.GET(walletID)
	if err == nil	{
		var balance float64
		balance, err = strconv.ParseFloat(cachedValue, 64)
		if err == nil	{
			return balance, nil
		}
		log.Println("Cache GET failed for walletID ", walletID, " ", err)
	}
	wallet, err := walletLogic.repo.FindWallet(walletID)
	if err != nil	{
		return 0, err
	}
	balance := wallet.Balance
	err = walletLogic.balanceCache.SET(walletID, balance)
	if err != nil	{
		log.Println("Cache SET failed for walletID ", walletID)
	}
	return balance, nil
}

func (walletLogic WalletLogicImpl) RequestWalletStatusChange(walletID int, status string)	string	{
	if !walletCommons.IsValidWalletStatus(status)	{
		return "Invalid Wallet Status!"
	}
	wallet, err := walletLogic.repo.FindWallet(walletID)
	if err == gorm.ErrRecordNotFound	{
		return errorhandling.HandleError(http.StatusBadRequest, "Wallet does not exist for given walletID", "Wallet does not exist for given walletID")
	}
	if err != nil	{
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	wallet.Status = status
	err = walletLogic.repo.SaveWallet(wallet)
	if err != nil	{
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	log.Println("Wallet status changed successfully")
	return "Wallet status changed successfully"
}

func (walletLogic WalletLogicImpl) DeleteWallet(walletID int) string {
	err := walletLogic.repo.DeleteWallet(walletID)
	if err != nil	{
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	log.Println("Wallet deleted successfully")
	return "Wallet deleted successfully"
}

func mapWalletRequestToWalletDao(walletRequestObject request.WalletRequest) *entities.Wallet {
	wallet := &entities.Wallet{
		UserId:  walletRequestObject.UserId,
		Balance: walletRequestObject.Balance,
		Status:  walletRequestObject.Status,
	}
	return wallet
}