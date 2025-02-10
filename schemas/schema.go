package schemas

import (
	"github.com/NochboolPrime/graphql-blog/resolvers"
	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(CommentType),
		},
		"allow_comments": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"post_id": &graphql.Field{
			Type: graphql.ID,
		},
		"parent_id": &graphql.Field{
			Type: graphql.ID,
		},
	},
})

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
			},
			Resolve: resolvers.GetPostResolver,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(PostType),
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"offset": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolvers.GetPostsResolver,
		},
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"allow_comments": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: resolvers.CreatePostResolver,
		},
		"createComment": &graphql.Field{
			Type: CommentType,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"parent_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
			},
			Resolve: resolvers.CreateCommentResolver,
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationType,
})
