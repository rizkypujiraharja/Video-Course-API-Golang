package entity

type OrderDetail struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	OrderID  uint   `gorm:"not null" json:"-"`
	LessonID string `gorm:"type:varchar(15)" json:"-"`
	Price    string `gorm:"type:varchar(15)" json:"-"`
	Lesson   Lesson `gorm:"foreignkey:LessonID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
