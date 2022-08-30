package logic

import (
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/api/dto/request"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/cache"
	transactionCommons "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/commons/transaction"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/commons/utils"
	walletCommons "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/commons/wallet"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/entities"
	errorhandling "bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/error-handling"
	"bitbucket.org/wallet_ledger_service/rahul_raj/wallet-as-a-service/repo"
	"encoding/csv"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TransactionLogic interface {
	RequestTransaction(transactionRequest *request.TransactionRequest) string
	SchedulePrevDayTransactionsCsvExport()
}

type TransactionLogicImpl struct {
	repo repo.Repo
	balanceCache cache.BalanceCache
}

func NewTransactionLogic(repository repo.Repo, cache cache.BalanceCache) TransactionLogic {
	return &TransactionLogicImpl{
		repo: repository,
		balanceCache: cache,
	}
}

func (transactionLogic TransactionLogicImpl) RequestTransaction(transactionRequest *request.TransactionRequest) string	{
	transaction := mapTransactionRequestToTransactionDao(transactionRequest)
	if transaction.Amount < 0 {
		return errorhandling.HandleError(http.StatusBadRequest, "Transaction failed: Negative transaction Amount invalid", "Transaction failed: Negative transaction Amount invalid")
	}

	wallet, err := transactionLogic.repo.FindWalletForUpdate(transaction.WalletId)
	if err == gorm.ErrRecordNotFound	{
		return errorhandling.HandleError(http.StatusBadRequest, "Wallet does not exist for given walletID", "Wallet does not exist for given walletID")
	}
	if err != nil{
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	if wallet.Status == walletCommons.BLOCKED.String() {
		return errorhandling.HandleError(http.StatusForbidden, "Transaction Denied: Wallet is blocked", "Transaction Denied: Wallet is blocked")
	}

	transactionSuccess, transactionErrorMessage := doTransaction(transaction, wallet)

	if transactionSuccess != true {
		return errorhandling.HandleError(http.StatusBadRequest, transactionErrorMessage, transactionErrorMessage)
	}

	err = transactionLogic.balanceCache.SET(transactionRequest.WalletId, wallet.Balance)
	if err != nil {
		log.Println("Failed to update cache" + err.Error())
	}

	err = transactionLogic.repo.SaveWallet(wallet)
	if err != nil {
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	err = transactionLogic.repo.SaveTransaction(transaction)
	if err != nil {
		return errorhandling.HandleError(http.StatusInternalServerError, "Something went wrong", err.Error())
	}
	log.Println("Transaction Successful")
	return "Transaction Successful"
}

func (transactionLogic TransactionLogicImpl) SchedulePrevDayTransactionsCsvExport() {
	cronJob := entities.CronJob{
		CronName:    "SchedulePrevDayTransactionsCsvExport",
		IsCompleted: false,
		RunDate: time.Now().Format("01-02-2006"),
	}
	rowsAffected, err := transactionLogic.repo.InsertIgnoreCronJob(&cronJob)
	if rowsAffected == 0	{
		log.Println("Cron already started by another instance")
		return
	}
	var transactions []entities.Transaction
	startDateTime := utils.GetYesterdayBeginning().Format("2006-01-02 15:04:05")
	endDateTime := utils.GetYesterdayEnd().Format("2006-01-02 15:04:05")
	transactions, err = transactionLogic.repo.FindTransactions(startDateTime, endDateTime)
	if err != nil {
		log.Println(err.Error())
		return
	}
	transactionsCSVRows := getTransactionCSVRows(transactions)
	err = utils.ExportToCsv(transactionsCSVRows, "transactions.csv")
	if err != nil {
		log.Println("Failed to export to CSV")
		return
	}
	log.Println("CSV Exported")
	cronJob.IsCompleted = true
	err = transactionLogic.repo.SaveCronStatus(&cronJob)
	if err != nil	{
		log.Println(err)
	}
}

func getTransactionCSVRows(transactions []entities.Transaction) [][]string {
	var rows [][]string
	headers := getTransactionHeaders()
	rows = append(rows, headers)
	for index, transaction := range transactions {
		var row []string
		row = append(row, strconv.Itoa(index))
		row = append(row, strconv.Itoa(transaction.WalletId))
		row = append(row, fmt.Sprint(transaction.Amount))
		row = append(row, transaction.TransactionType)
		rows = append(rows, row)
	}
	return rows
}

func getTransactionHeaders() []string {
	var headers []string
	headers = append(headers, "SNo.")
	headers = append(headers, "Wallet ID")
	headers = append(headers, "Amount")
	headers = append(headers, "Transaction Type")
	return headers
}


func doTransaction(transaction *entities.Transaction, wallet *entities.Wallet) (bool, string) {
	var transactionSuccess bool
	var transactionErrorMessage string
	switch transaction.TransactionType {
	case transactionCommons.CREDIT.String():
		transactionSuccess, transactionErrorMessage = credit(wallet, transaction)
	case transactionCommons.DEBIT.String():
		transactionSuccess, transactionErrorMessage = debit(wallet, transaction)
	default:
		transactionSuccess, transactionErrorMessage = false, "Invalid transaction type"
	}
	return transactionSuccess, transactionErrorMessage
}

func mapTransactionRequestToTransactionDao(transactionRequest *request.TransactionRequest) *entities.Transaction {
	transaction := &entities.Transaction{
		WalletId:        transactionRequest.WalletId,
		Amount:          transactionRequest.Amount,
		TransactionType: transactionRequest.TransactionType,
	}
	return transaction
}

func debit(wallet *entities.Wallet, transaction *entities.Transaction) (bool, string) {
	wallet.Balance -= transaction.Amount
	if wallet.Balance < 0 {
		return false,"Transaction failed, this transaction will result in a lower that threshold balance of " + string(rune(walletCommons.MinWalletBalance))
	}
	return true,""
}

func credit(wallet *entities.Wallet, transaction *entities.Transaction) (bool, string){
	wallet.Balance += transaction.Amount
	if wallet.Balance > walletCommons.MaxWalletBalance {
		return false,"Transaction failed, this transaction will result in a higher than upper limit balance of " + string(rune(walletCommons.MaxWalletBalance))
	}
	return true,""
}

package main

import (
"encoding/csv"
"fmt"
"os"
)

func main() {
	masterFile := "user_credit_limit_hierarchy_202104281407.csv"
	otherFile := "user_credit_limit_hierarchy_202104291125.csv"
	masterRecords := csvReaderAll(masterFile)
	otherRecords := csvReaderAll(otherFile)
	var commonRows [][]string

	for _, masterRow := range masterRecords	{
		for _, otherRow := range otherRecords 	{
			if masterRow[1] == otherRow[1] && masterRow[9] == otherRow[9]	{
				commonRow := []string{masterRow[1], masterRow[9]}
				commonRows = append(commonRows, commonRow)
			}
		}
	}

	exportToCsv(commonRows, "common.csv")
}

func csvReaderAll(fileName string) [][]string {
	// Open the file
	recordFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil
	}
	// fmt.Println(len(allRecords))

	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil
	}

	return allRecords
}

func exportToCsv(data [][]string, csvName string) error {
	file, err := os.Create(csvName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
}