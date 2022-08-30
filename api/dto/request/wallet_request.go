package request

type WalletRequest struct {
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
}
