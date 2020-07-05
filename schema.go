package main

import (
	"graphql-ca/src/graphql/posts/serializer"

	"github.com/graphql-go/graphql"
)

type Schema struct {
	resolver Resolver
}

// NewSchema initializes Schema struct which takes resolver as the argument.
func NewSchema() Schema {
	resolver := NewResolver()
	return Schema{
		resolver: resolver,
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
				Resolve: s.resolver.GetPosts,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}
