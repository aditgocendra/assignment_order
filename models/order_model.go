package models

import (
	"time"
)

type Orders struct {
	OrderID      uint 		`gorm:"primaryKey"`
	CustomerName string		`gorm:"not null;type:varchar(191)"`
	Items 		 []Items 	`gorm:"foreignKey:OrderID"`
	OrderedAt    time.Time
}