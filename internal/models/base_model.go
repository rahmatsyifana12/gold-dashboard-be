package models

import "time"

type BaseModel struct {
	ID        uint           `gorm:"id,primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"-"`
}
