package models

import (
	"gorm.io/gorm"
)

type WishStatus string

// WishStatus represents the status of a wish.
// @Description Status of the wish
// @Enum approved
// @Enum pending
// @Enum rejected
const (
	Approved WishStatus = "approved"
	Pending  WishStatus = "pending"
	Rejected WishStatus = "rejected"
)

// Wish represents a wish object in the system.
// @Description Wish object
// @ID wish
// @Property title string "Title of the wish" example("A wish")
// @Property content string "Content of the wish" example("This is the content")
// @Property category string "Category of the wish" example("General")
// @Property is_confirm string "Status of the wish" example("approved")
// @Property is_deleted bool "Indicates if the wish is deleted" example(false)
type Wish struct {
	gorm.Model
	Title     string     `json:"title" gorm:"size:100"`
	Content   string     `json:"content" gorm:"size:255"`
	Category  string     `json:"category" gorm:"size:50"`
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
