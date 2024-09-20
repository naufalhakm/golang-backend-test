package models

import "time"

type Author struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	Birthdate time.Time `gorm:"type:date"`
}
