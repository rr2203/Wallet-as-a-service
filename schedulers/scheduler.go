package schedulers

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/logic"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"github.com/mileusna/crontab"
	"log"
)

var transactionLogic = logic.NewTransactionLogic(repo.NewMySQLRepo(), cache.NewRedisCache("localhost:6379", 0, 0))

func InitializeCronJobs() {
	Schedule(transactionLogic.SchedulePrevDayTransactionsCsvExport, "* * * * *")
}

func Schedule(job func(), schedule string) {
	err := crontab.New().AddJob(schedule, job)
	if err != nil {
		log.Println(err)
		return
	}
}