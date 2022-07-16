package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/boodyvo/testswapi/models"
)

const (
	allMoviesAPIURL = "https://swapi.dev/api/films"
	movieAPIURL     = "https://swapi.dev/api/films/%s"
)

func fetchMovie(ctx context.Context, id string) (*moviesResponse, error) {
	url := fmt.Sprintf(movieAPIURL, id)
	body, err := fetch(ctx, url)
	if err != nil {
		return nil, err
	}

	var data *moviesResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func fetchCharacter(ctx context.Context, urls []string) ([]*models.Character, error) {
	characters := make([]*models.Character, 0, len(urls))
	g, subCtx := errgroup.WithContext(ctx)

	for _, url := range urls {
		characterURL := url
		g.Go(func() error {
			body, err := fetch(subCtx, characterURL)
			if err != nil {
				return err
			}

			var character *characterResponse
			if err := json.Unmarshal(body, &character); err != nil {
				return err
			}

			characters = append(characters, character.ToCharacter())

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return characters, nil
}

func fetchMovies(ctx context.Context) ([]*models.Movie, error) {
	body, err := fetch(ctx, allMoviesAPIURL)
	if err != nil {
		return nil, err
	}

	var data response
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	movies := make([]*models.Movie, 0, len(data.Results))
	for _, movie := range data.Results {
		movies = append(movies, movie.ToMovie())
	}

	return movies, nil
}

func fetch(_ context.Context, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return body, err
}
