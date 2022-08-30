package transaction

type TransactionType int
const (
	TRANSACTION_TYPE_INVALID TransactionType = iota
	CREDIT
	DEBIT
)
var TransactionTypeToStringMap = map[TransactionType]string{
	CREDIT: "CREDIT",
	DEBIT:  "DEBIT",
}
var StringToTransactionTypeMap = map[string]TransactionType{
	"CREDIT": CREDIT,
	"DEBIT":  DEBIT,
}
func (transactionType TransactionType) String() string {
	return TransactionTypeToStringMap[transactionType]
}
func IsValidTransactionType(transactionType string) bool {
	_, exist := StringToTransactionTypeMap[transactionType]
	return exist
}