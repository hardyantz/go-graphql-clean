package repository

import (
	"database/sql"
	"graphql-ca/config"
	postDomain "graphql-ca/src/graphql/posts/domain"
)

type posts struct {
	DB *sql.DB
}

func NewPostsRepository(conf config.Config) *posts {
	return &posts{conf.DB}
}

func (p *posts) FetchAll(search string) ([]postDomain.Posts, error) {
	var result []postDomain.Posts
	rows, err := p.DB.Query("SELECT title, description FROM posts where title like '%" + search + "%'")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post postDomain.Posts
		err = rows.Scan(&post.Title, &post.Description)
		if err != nil {
			return result, err
		}

		result = append(result, post)
	}

	return result, nil
}
