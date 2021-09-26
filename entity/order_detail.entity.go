package entity

type OrderDetail struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	OrderID  int64  `gorm:"not null" json:"-"`
	LessonID int64  `gorm:"not null" json:"-"`
	Price    int64  `gorm:"type:bigint" json:"-"`
	Lesson   Lesson `gorm:"foreignkey:LessonID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
