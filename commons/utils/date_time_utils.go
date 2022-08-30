package utils

import "time"

func GetYesterday() (int, time.Month, int) {
	return time.Now().AddDate(0, 0, -1).Date()
}

func GetYesterdayEnd() time.Time {
	yesterdayYear, yesterdayMonth, yesterdayDay := GetYesterday()
	return time.Date(yesterdayYear, yesterdayMonth, yesterdayDay, 23, 59, 59, 0, time.UTC)
}

func GetYesterdayBeginning() time.Time {
	yesterdayYear, yesterdayMonth, yesterdayDay := GetYesterday()
	return time.Date(yesterdayYear, yesterdayMonth, yesterdayDay, 0, 0, 0, 0, time.UTC)
}