package repo

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DBInstance *gorm.DB
var err error
const DNS = "root:rohitraj1@tcp(127.0.0.1:3306)/godb?parseTime=true"

func (MySQLRepo) SetupDB()	{
	DBInstance, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil	{
		log.Println(err.Error())
		panic("Cannot connect to DBInstance")
	}

	err := DBInstance.AutoMigrate(&entities.User{})
	if err != nil {
		log.Println(err.Error())
		panic("Automigration failed for User")
	}
	err = DBInstance.AutoMigrate(&entities.Wallet{})
	if err != nil {
		log.Println(err.Error())
		panic("Automigration failed for Wallet")
	}
	err = DBInstance.AutoMigrate(&entities.Transaction{})
	if err != nil {
		log.Println(err.Error())
		panic("Automigration failed for Transaction")
	}

	err = DBInstance.AutoMigrate(&entities.CronJob{})
	if err != nil {
		log.Println(err.Error())
		panic("Automigration failed for CronJob")
	}
}