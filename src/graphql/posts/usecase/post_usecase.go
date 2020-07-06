package usecase

import (
	postDomain "graphql-ca/src/graphql/posts/domain"
	postRepository "graphql-ca/src/graphql/posts/repository"
)

type PostUsecase interface {
	GetAll(search string) ([]postDomain.Posts, error)
}

type postUsecase struct {
	repo postRepository.Posts
}

func NewPostUseCase(repo postRepository.Posts) *postUsecase {
	return &postUsecase{repo: repo}
}

func (p *postUsecase) GetAll(search string) ([]postDomain.Posts, error) {
	res, err := p.repo.FetchAll(search)

	return res, err
}
