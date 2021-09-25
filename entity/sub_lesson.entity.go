package entity

type SubLesson struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	Title    string `gorm:"type:varchar(255)" json:"-"`
	LessonID uint   `gorm:"not null" json:"-"`
}
