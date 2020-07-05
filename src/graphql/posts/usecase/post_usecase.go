package usecase

import (
	postDomain "graphql-ca/src/graphql/posts/domain"
	postRepository "graphql-ca/src/graphql/posts/repository"
)

type PostUseCaseImplementation interface {
	GetAll(search string) ([]postDomain.Posts, error)
}

type PostUseCase struct {
	repo postRepository.PostImplementation
}

func NewPostUseCase(repo postRepository.PostImplementation) *PostUseCase {
	return &PostUseCase{repo: repo}
}

func (p *PostUseCase) GetAll(search string) ([]postDomain.Posts, error) {
	res, err := p.repo.FetchAll(search)

	return res, err
}
