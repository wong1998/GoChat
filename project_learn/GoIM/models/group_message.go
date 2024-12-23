package models

type GroupMessage struct {
	ID        uint   `gorm:"primaryKey"`
	GroupID   uint   `gorm:"not null"`
	SenderID  uint   `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	Timestamp int64  `gorm:"autoCreateTime"`
}
