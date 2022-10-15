package models

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"type:varchar(200); not null" json:"customerName"`
	OrderedAt    time.Time `gorm:"default: NOW()" json:"orderedAt"`
	Items        []Item
}
