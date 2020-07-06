package repository

import postDomain "graphql-ca/src/graphql/posts/domain"

type Posts interface {
	FetchAll(search string) ([]postDomain.Posts, error)
}
