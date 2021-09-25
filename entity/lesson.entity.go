package entity

type Lesson struct {
	ID              int64       `gorm:"primary_key:auto_increment" json:"-"`
	Title           string      `gorm:"type:varchar(255)" json:"-"`
	Description     string      `gorm:"type:text" json:"-"`
	Price           uint64      `gorm:"type:bigint" json:"-"`
	ImageCoverUrl   string      `gorm:"type:varchar(255)" json:"-"`
	VideoPreviewUrl string      `gorm:"type:varchar(255)" json:"-"`
	CategoryID      int64       `gorm:"not null" json:"-"`
	Category        Category    `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	SubLessons      []SubLesson `json:"-"`
}
