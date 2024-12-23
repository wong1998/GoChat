package models

type Group struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
