package main

import (
	postResolver "graphql-ca/src/graphql/posts/resolver"
	"graphql-ca/src/graphql/posts/serializer"

	"github.com/graphql-go/graphql"
)

type Schema struct {
	postResolver postResolver.PostResolver
}

// NewSchema initializes Schema struct which takes resolver as the argument.
func NewSchema() Schema {
	postResolverImpl := postResolver.NewPostResolver()
	return Schema{
		postResolver: postResolverImpl,
	}
}

// Query initializes config schema query for graphql server.
func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Posts": &graphql.Field{
				Type:        serializer.SearchResults,
				Description: "Fetch Search Posts",
				Args: graphql.FieldConfigArgument{
					"search": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.postResolver.GetPosts,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}
