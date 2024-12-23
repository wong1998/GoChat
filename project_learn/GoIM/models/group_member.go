package models

type GroupMember struct {
	ID      uint `gorm:"primaryKey"`
	GroupID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
}
