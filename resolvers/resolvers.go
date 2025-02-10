package resolvers

import (
	"log"
	"strconv"

	"github.com/NochboolPrime/graphql-blog/database"
	"github.com/NochboolPrime/graphql-blog/models"
	"github.com/graphql-go/graphql"
)

func GetPostResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if ok {
		var post models.Post
		if err := database.DB.Preload("Comments").First(&post, id).Error; err != nil {
			log.Printf("Error retrieving post: %v", err)
			return nil, err
		}
		return post, nil
	}
	return nil, nil
}

func GetPostsResolver(p graphql.ResolveParams) (interface{}, error) {
	var posts []models.Post

	limit, ok := p.Args["limit"].(int)
	if !ok {
		limit = 10
	}
	offset, ok := p.Args["offset"].(int)
	if !ok {
		offset = 0
	}

	if err := database.DB.Preload("Comments").Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		log.Printf("Error retrieving posts: %v", err)
		return nil, err
	}
	return posts, nil
}

func CreatePostResolver(p graphql.ResolveParams) (interface{}, error) {
	title := p.Args["title"].(string)
	content := p.Args["content"].(string)
	allowComments := p.Args["allow_comments"].(bool)

	post := models.Post{
		Title:         title,
		Content:       content,
		AllowComments: allowComments,
	}
	if err := database.DB.Create(&post).Error; err != nil {
		log.Printf("Error creating post: %v", err)
		return nil, err
	}
	return post, nil
}

func CreateCommentResolver(p graphql.ResolveParams) (interface{}, error) {
	postIDStr := p.Args["post_id"].(string)
	text := p.Args["text"].(string)
	parentIDStr, _ := p.Args["parent_id"].(string)

	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		log.Printf("Error parsing post ID: %v", err)
		return nil, err
	}

	var parentID uint
	if parentIDStr != "" {
		parsedParentID, err := strconv.ParseUint(parentIDStr, 10, 32)
		if err != nil {
			log.Printf("Error parsing parent ID: %v", err)
			return nil, err
		}
		parentID = uint(parsedParentID)
	}

	comment := models.Comment{
		PostID:   uint(postID),
		Text:     text,
		ParentID: parentID,
	}
	if err := database.DB.Create(&comment).Error; err != nil {
		log.Printf("Error creating comment: %v", err)
		return nil, err
	}
	return comment, nil
}
