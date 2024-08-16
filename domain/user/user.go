package user

import "time"

// User represents the user entity.
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      *string   `gorm:"size:255"`
	Email     string    `gorm:"unique;size:255"`
	Password  string    `gorm:"size:255"`
	Image     *string   `gorm:"size:255"`
	Role      string    `gorm:"default:member;size:50"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
