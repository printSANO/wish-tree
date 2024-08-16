package models

import (
	"gorm.io/gorm"
)

type WishStatus string

const (
	Approved WishStatus = "approved"
	Pending  WishStatus = "pending"
	Rejected WishStatus = "rejected"
)

// gorm.model에는 ID, CreatedAt, UpdatedAt, DeletedAt이 포함되어 있음
type Wish struct {
	gorm.Model
	Title     string     `json:"title" gorm:"size:100"`
	Content   string     `json:"content" gorm:"size:255"`
	Category  string     `json:"category" gorm:"size:50"`
	IsConfirm WishStatus `json:"is_confirm" gorm:"type:wish_status"`
	IsDeleted bool       `json:"is_deleted"`
}
