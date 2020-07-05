package repository

import postDomain "graphql-ca/src/graphql/posts/domain"

type PostImplementation interface {
	FetchAll(search string) ([]postDomain.Posts, error)
}
