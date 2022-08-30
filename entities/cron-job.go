package entities

import (
	"gorm.io/gorm"
)

type CronJob struct {
	gorm.Model
	CronName string `gorm:"index:idx_name,unique"`
	IsCompleted bool
	RunDate string `gorm:"index:idx_name,unique"`
}