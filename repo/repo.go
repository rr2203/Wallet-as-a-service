package repo

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
)

type Repo interface {
	SetupDB()

	CreateUser(user *entities.User) error
	GetUsers() (*[]entities.User, error)
	GetUser(userId int) (*entities.User, error)
	UpdateUser(user *entities.User, userID int) error
	DeleteUser(userID int) error

	CreateWallet(wallet *entities.Wallet) error
	FindWallet(walletID int) (*entities.Wallet, error)
	FindWalletForUpdate(walletID int) (*entities.Wallet, error)
	SaveWallet(wallet *entities.Wallet) error
	DeleteWallet(walletID int) error

	SaveTransaction(transaction *entities.Transaction) error
	FindTransactions(startDateTime string, endDateTime string) ([]entities.Transaction, error)

	InsertIgnoreCronJob(cronJob *entities.CronJob) (int64, error)
	SaveCronStatus(cronStatus *entities.CronJob) error
}