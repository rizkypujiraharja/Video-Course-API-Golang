package entity

type Category struct {
	ID   int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name string `gorm:"type:varchar(100);unique;" json:"-"`
	Slug string `gorm:"type:varchar(100);unique;" json:"-"`
}
