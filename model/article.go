package model

import (
	"gorm.io/gorm"
)

type Article struct {
	category Category
	gorm.Model
	Title   string `gorm:"type:varchar(100); not null" json:"title"`
	Cid     int    `gorm:"type:int; not null" json:"cid"`
	Desc    string `gorm:"type:varchar(20); not null" json:"desc"`
	Content string `gorm:"type:varchar(100); not null" json:"content"`
	Img     string `gorm:"type:varchar(100); null" json:"img"`
}
