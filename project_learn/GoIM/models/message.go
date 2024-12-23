package models

type Message struct {
	ID         uint   `gorm:"primaryKey"`
	SenderID   uint   `gorm:"not null"`
	ReceiverID uint   `gorm:"not null"`
	Content    string `gorm:"type:text;not null"`
	Timestamp  int64  `gorm:"autoCreateTime"`
}
