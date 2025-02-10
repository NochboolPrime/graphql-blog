package model

type Post struct {
	ID            string    `json:"id" gorm:"primary_key"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Comments      []Comment `json:"comments" gorm:"foreignKey:PostID"`
	AllowComments bool      `json:"allowComments"`
}

type Comment struct {
	ID       string `json:"id" gorm:"primary_key"`
	Text     string `json:"text"`
	PostID   string `json:"postID"`
	ParentID string `json:"parentID"`
}
