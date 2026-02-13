// Copyright by AcmaTvirus
package models

import (
	"time"
)

// Project đại diện cho một dự án Docker của người dùng
type Project struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Image     string    `json:"image"`
	Status    string    `json:"status"` // running, stopped, etc.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
