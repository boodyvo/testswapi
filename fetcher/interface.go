package fetcher

import (
	"context"
	"github.com/boodyvo/testswapi/models"
)

type CharacterFilter struct {
	Gender string
}

type CharacterSort struct {
	IsAsc    bool
	SortName string
}

type CharacterRequest struct {
	MovieID string
	Filter  *CharacterFilter
	Sort    *CharacterSort
}

type StarWarsAPI interface {
	ListMovies(ctx context.Context) ([]*models.Movie, error)
	ListCharacters(ctx context.Context, request CharacterRequest) ([]*models.Character, int, int, error)
}
