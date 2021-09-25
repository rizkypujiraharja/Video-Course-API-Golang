package entity

type Video struct {
	ID          int64     `gorm:"primary_key:auto_increment" json:"-"`
	Title       string    `gorm:"type:varchar(255)" json:"-"`
	Description string    `gorm:"type:text" json:"-"`
	VideoUrl    string    `gorm:"type:varchar(255)" json:"-"`
	SubLessonID int64     `gorm:"not null" json:"-"`
	SubLesson   SubLesson `gorm:"foreignkey:SubLessonID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
