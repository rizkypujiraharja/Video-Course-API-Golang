package entity

type OrderedLesson struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	UserID   int64  `gorm:"not null" json:"-"`
	LessonID int64  `gorm:"not null" json:"-"`
	Lesson   Lesson `gorm:"foreignkey:LessonID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	User     User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
