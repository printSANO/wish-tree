package models

import (
	"gorm.io/gorm"
)

// WishStatus represents the status of a wish.
// @Description Status of the wish
// @Enum approved
// @Enum pending
// @Enum rejected
type WishStatus string

const (
	Approved WishStatus = "approved"
	Pending  WishStatus = "pending"
	Rejected WishStatus = "rejected"
)

// Wish represents a wish object in the system.
// @Description Wish object containing details of a user's wish.
// @ID wish
type Wish struct {
	// The unique ID of the wish.
	// @Description The auto-incrementing ID of the wish.
	// @example 1
	gorm.Model

	// The title of the wish.
	// @Description Title of the wish.
	// @example A wish
	Title string `json:"title" gorm:"size:100"`

	// The content or detailed description of the wish.
	// @Description Detailed content of the wish.
	// @example This is the content of my wish.
	Content string `json:"content" gorm:"size:255"`

	// The category of the wish, which groups similar wishes together.
	// @Description Category of the wish.
	// @example General
	Category string `json:"category" gorm:"size:50"`

	// The current status of the wish (approved, pending, rejected).
	// @Description Status of the wish.
	// @Enum approved
	// @Enum pending
	// @Enum rejected
	// @example pending
	IsConfirm WishStatus `json:"is_confirm" gorm:"type:wish_status;default:'pending'"`
}

type Filter struct {
	Category string `json:"category" form:"category"`
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	Page     int    `json:"page" form:"page"`
	Limit    int    `json:"limit" form:"limit"`
}

type Pagination struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type PaginatedResponse[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}
