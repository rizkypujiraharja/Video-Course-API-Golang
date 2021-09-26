package entity

import "time"

type Order struct {
	ID           int64         `gorm:"primary_key:auto_increment" json:"-"`
	UserID       int64         `gorm:"not null" json:"-"`
	InvoiceCode  string        `gorm:"type:varchar(15)" json:"-"`
	Status       string        `gorm:"type:varchar(15)" json:"-"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	User         User          `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	OrderDetails []OrderDetail `json:"-"`
}
