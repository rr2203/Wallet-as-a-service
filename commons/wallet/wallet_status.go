package wallet

type WalletStatus int
const (
	WALLET_STATUS_INVALID WalletStatus = iota
	ACTIVE
	BLOCKED
)
var WalletStatusToStringMap = map[WalletStatus]string{
	ACTIVE:  "ACTIVE",
	BLOCKED: "BLOCKED",
}
var StringToWalletStatusMap = map[string]WalletStatus{
	"ACTIVE":  ACTIVE,
	"BLOCKED": BLOCKED,
}
func (walletStatus WalletStatus) String() string {
	return WalletStatusToStringMap[walletStatus]
}
func IsValidWalletStatus(walletStatus string) bool {
	_, exist := StringToWalletStatusMap[walletStatus]
	return exist
}