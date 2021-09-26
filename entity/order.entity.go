package entity

import "time"

type Order struct {
	ID           int64         `gorm:"primary_key:auto_increment" json:"-"`
	UserID       int64         `gorm:"not null" json:"-"`
	InvoiceCode  string        `gorm:"type:varchar(45);unique;" json:"-"`
	Status       string        `gorm:"type:varchar(15)" json:"-"`
	TotalOrder   uint64        `gorm:"type:bigint" json:"-"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	User         User          `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	OrderDetails []OrderDetail `json:"-"`
}
