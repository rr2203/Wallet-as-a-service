package repo

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) SetupDB() {
}

func (mock *MockRepository) CreateUser(user *entities.User) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) GetUsers() (*[]entities.User, error)	{
	args := mock.Called()
	return args.Get(0).(*[]entities.User), args.Error(1)
}

func (mock *MockRepository) GetUser(userId int) (*entities.User, error)	{
	args := mock.Called()
	return args.Get(0).(*entities.User), args.Error(1)
}

func (mock *MockRepository) UpdateUser(user *entities.User, userID int) error	{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) DeleteUser(userID int) error	{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) CreateWallet(wallet *entities.Wallet) error	{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) FindWallet(walletID int) (*entities.Wallet, error)	{
	args := mock.Called()
	return args.Get(0).(*entities.Wallet), args.Error(1)
}

func (mock *MockRepository) FindWalletForUpdate(walletID int) (*entities.Wallet, error)	{
	args := mock.Called()
	return args.Get(0).(*entities.Wallet), args.Error(1)
}

func (mock *MockRepository) SaveWallet(wallet *entities.Wallet) error{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) DeleteWallet(walletID int) error{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) SaveTransaction(transaction *entities.Transaction) error{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) FindTransactions(startDateTime string, endDateTime string) ([]entities.Transaction, error) {
	args := mock.Called()
	return args.Get(0).([]entities.Transaction), args.Error(1)
}

func (mock *MockRepository) FindPendingCronForUpdate(cronName string, runDate string) (*entities.CronJob, error) {
	args := mock.Called()
	return args.Get(0).(*entities.CronJob), args.Error(1)
}

func (mock *MockRepository) SaveCronStatus(cronStatus *entities.CronJob) error{
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) InsertIgnoreCronJob(cronJob *entities.CronJob) (int64, error) {
	args := mock.Called()
	return args.Get(0).(int64), args.Error(1)
}