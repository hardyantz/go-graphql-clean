package resolver

import (
	"errors"
	"graphql-ca/config"
	"graphql-ca/helper"
	postRepository "graphql-ca/src/graphql/posts/repository"
	"graphql-ca/src/graphql/posts/usecase"
	"net/http"

	"github.com/graphql-go/graphql"
)

type postResolver struct {
	postUsecase usecase.PostUsecase
}

type PostResolver interface {
	GetPosts(params graphql.ResolveParams) (interface{}, error)
}

func NewPostResolver() *postResolver{
	cfg := config.InitConfig()

	repo := postRepository.NewPostsRepository(cfg)
	postUseCase := usecase.NewPostUseCase(repo)

	return &postResolver{postUseCase}
}

func (r *postResolver) GetPosts(params graphql.ResolveParams) (interface{}, error) {

	postSearch, postSearchOk := params.Args["search"].(string)
	if !postSearchOk {
		return helper.NewHTTPResponse(http.StatusBadRequest, "invalid params", nil), errors.New("invalid params")
	}

	res, err := r.postUsecase.GetAll(postSearch)
	if err != nil {
		return helper.NewHTTPResponse(http.StatusBadRequest, err.Error(), nil), err
	}

	return helper.NewHTTPResponse(http.StatusOK, "get posts", res), nil
}
