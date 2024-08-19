package models

import "gorm.io/gorm"

// Comment represents a comment on a wish.
// @Description A comment that belongs to a wish.
type Comment struct {
	// The unique ID of the comment.
	// @Description The auto-incrementing ID of the comment.
	// @example 1
	gorm.Model

	// The content of the comment.
	// @Description The actual text of the comment.
	// @example This is a great wish!
	Content string `json:"content"`

	// The ID of the wish this comment belongs to.
	// @Description The ID of the associated wish.
	// @example 5
	WishID uint `json:"wish_id"`

	// The associated wish.
	// @Description The Wish object that this comment is associated with.
	Wish Wish `json:"wish" gorm:"foreignKey:WishID"`
}

type CommentLimit struct {
	// The number of comments to return.
	// @Description The number of comments to return.
	// @example 10
	Limit int `json:"limit" form:"limit"`

	// The page number of comments to return.
	// @Description The page number of comments to return.
	// @example 1
	Page int `json:"page" form:"page"`
}
