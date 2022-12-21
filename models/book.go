package models

type Book struct {
	Id          int64  `gorm:"primarykey" json:"id"`
	Tittle      string `gorm:"type:varchar(300)" json:"tittle"`
	Description string `gorm:"type:text" json:"description"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	PublishDate string `gorm:"type:date" json:"publishdate"`
}
