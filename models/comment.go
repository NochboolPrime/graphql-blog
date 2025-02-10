package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	PostID   uint   `json:"post_id"`
	Text     string `json:"text"`
	ParentID uint   `json:"parent_id,omitempty"`
}
