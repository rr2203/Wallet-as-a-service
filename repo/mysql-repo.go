package repo

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MySQLRepo struct {
}

func NewMySQLRepo() Repo {
	return &MySQLRepo{}
}

func (MySQLRepo) CreateUser(user *entities.User) error {
	result := DBInstance.Create(user)
	return result.Error
}

func (MySQLRepo) GetUsers() (*[]entities.User, error)	{
	var users []entities.User
	result := DBInstance.Find(&users)
	return &users, result.Error
}

func (MySQLRepo) GetUser(userId int) (*entities.User, error)	{
	var user *entities.User
	result := DBInstance.First(&user, userId)
	return user, result.Error
}

func (MySQLRepo) UpdateUser(user *entities.User, userID int) error	{
	var prevUser entities.User
	if result := DBInstance.First(&prevUser, userID); result.Error != nil	{
		return result.Error
	}
	prevUser.Email = user.Email
	prevUser.FirstName = user.FirstName
	prevUser.LastName = user.LastName
	result := DBInstance.Save(&prevUser)
	return result.Error
}

func (MySQLRepo) DeleteUser(userID int) error	{
	var user []entities.User
	result := DBInstance.Delete(&user, userID)
	return result.Error
}

func (MySQLRepo) CreateWallet(wallet *entities.Wallet) error	{
	result := DBInstance.Create(wallet)
	return result.Error
}

func (MySQLRepo) FindWallet(walletID int) (*entities.Wallet, error)	{
	var wallet entities.Wallet
	result := DBInstance.Find(&wallet, walletID)
	if (wallet == entities.Wallet{}) {
		return nil, gorm.ErrRecordNotFound
	}
	return &wallet, result.Error
}

func (MySQLRepo) FindWalletForUpdate(walletID int) (*entities.Wallet, error)	{
	var wallet entities.Wallet
	result := DBInstance.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&wallet, walletID)
	if (wallet == entities.Wallet{}) {
		return nil, gorm.ErrRecordNotFound
	}
	return &wallet, result.Error
}

func (MySQLRepo) SaveWallet(wallet *entities.Wallet) error{
	result := DBInstance.Save(wallet)
	return result.Error
}

func (MySQLRepo) DeleteWallet(walletID int) error{
	var wallet entities.Wallet
	result := DBInstance.Find(&wallet, walletID)
	if result.Error != nil	{
		return result.Error
	}
	result = DBInstance.Delete(&wallet)
	return result.Error
}

func (MySQLRepo) SaveTransaction(transaction *entities.Transaction) error{
	result := DBInstance.Save(transaction)
	return result.Error
}

func (MySQLRepo) FindTransactions(startDateTime string, endDateTime string) ([]entities.Transaction, error) {
	var transaction []entities.Transaction
	result := DBInstance.Where("created_at BETWEEN ? AND ?", startDateTime, endDateTime).Find(&transaction)
	return transaction, result.Error
}

func (MySQLRepo) InsertIgnoreCronJob(cronJob *entities.CronJob) (int64, error) {
	result := DBInstance.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&cronJob)
	return result.RowsAffected, result.Error
}

func (MySQLRepo) SaveCronStatus(cronStatus *entities.CronJob) error{
	result := DBInstance.Save(cronStatus)
	return result.Error
}