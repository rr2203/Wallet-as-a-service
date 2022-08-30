package request

type StatusChangeRequest struct {
	WalletID int    `json:"wallet_id"`
	Status   string `json:"status"`
}
