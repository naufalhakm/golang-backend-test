package models

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:255"`
	ISBN     string `gorm:"unique"`
	AuthorID uint
	Author   Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
