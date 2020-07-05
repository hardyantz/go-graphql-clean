package serializer

import "github.com/graphql-go/graphql"

var SearchResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResult",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type: graphql.Boolean,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: graphql.NewList(SearchPosts),
			},
		},
	},
)

var SearchPosts = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchPosts",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)