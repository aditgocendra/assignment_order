package models

type Items struct {
	ItemID      uint   `gorm:"primaryKey"`
	ItemCode    int    `gorm:"not null;type:int"`
	Description string `gorm:"not null;type:varchar(191)"`
	OrderID     uint
	Quantity    int `gorm:"not null;type:int"`
}