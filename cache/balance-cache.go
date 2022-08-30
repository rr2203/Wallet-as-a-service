package cache

type BalanceCache interface {
	SET(key int, value float64) error
	GET(key int) (string, error)
}