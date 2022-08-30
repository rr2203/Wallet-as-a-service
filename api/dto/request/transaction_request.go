package request

type TransactionRequest struct {
	WalletId  int  `json:"wallet_id"`
	Amount float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
}