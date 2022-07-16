package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/boodyvo/testswapi/fetcher"
	"github.com/boodyvo/testswapi/models"
	"github.com/go-redis/redis/v8"
)

const (
	layout = "2006-01-02"
)

type Fetcher struct {
	client *redis.Client
}

func New(url string) (*Fetcher, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0, // use default DB
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &Fetcher{client: client}, nil
}

// ListMovies list movies sorted by release_date ascending
func (f *Fetcher) ListMovies(ctx context.Context) ([]*models.Movie, error) {
	var movies []*models.Movie
	val, err := f.client.Get(ctx, "all_movies").Result()
	if err == redis.Nil {
		return f.fetchAndSaveMovies(ctx)
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &movies)
	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ReleaseDate.Before(movies[j].ReleaseDate)
	})

	return movies, err

}

func (f *Fetcher) fetchAndSaveMovies(ctx context.Context) ([]*models.Movie, error) {
	movies, err := fetchMovies(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ReleaseDate.Before(movies[j].ReleaseDate)
	})

	moviesBytes, err := json.Marshal(movies)
	if err != nil {
		return nil, err
	}

	if err := f.client.Set(ctx, "all_movies", moviesBytes, 0).Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// ListCharacters list characters for a particular movie sorted and filtered based on request.
// Returns list of characters, count of total filtered characters and total height
func (f *Fetcher) ListCharacters(ctx context.Context, request fetcher.CharacterRequest) ([]*models.Character, int, int, error) {
	var characters []*models.Character
	key := fmt.Sprintf("characters_%s", request.MovieID)
	val, err := f.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return f.fetchAndSaveCharacters(ctx, request)
	}
	if err != nil {
		return nil, 0, 0, err
	}

	err = json.Unmarshal([]byte(val), &characters)
	if err != nil {
		return nil, 0, 0, err
	}

	return processCharacters(ctx, request, characters)
}

func (f *Fetcher) fetchAndSaveCharacters(ctx context.Context, request fetcher.CharacterRequest) ([]*models.Character, int, int, error) {
	key := fmt.Sprintf("characters_%s", request.MovieID)

	movie, err := fetchMovie(ctx, request.MovieID)
	if err != nil {
		return nil, 0, 0, err
	}

	characters, err := fetchCharacter(ctx, movie.Characters)
	if err != nil {
		return nil, 0, 0, err
	}

	charactersBytes, err := json.Marshal(characters)
	if err != nil {
		return nil, 0, 0, err
	}

	if err := f.client.Set(ctx, key, charactersBytes, 0).Err(); err != nil {
		return nil, 0, 0, err
	}

	return processCharacters(ctx, request, characters)
}

func processCharacters(_ context.Context, request fetcher.CharacterRequest, characters []*models.Character) ([]*models.Character, int, int, error) {
	totalCount := 0
	totalHeight := 0
	results := make([]*models.Character, 0, len(characters))
	for _, character := range characters {
		if request.Filter != nil &&
			character.Gender != request.Filter.Gender {
			continue
		}

		results = append(results, character)
		totalCount++
		totalHeight += character.Height
	}

	if request.Sort != nil {
		switch request.Sort.SortName {
		case "name":
			sort.Slice(results, func(i, j int) bool {
				if request.Sort.IsAsc {
					return results[i].Name < results[j].Name
				}

				return results[i].Name > results[j].Name
			})
		case "gender":
			sort.Slice(results, func(i, j int) bool {
				if request.Sort.IsAsc {
					return results[i].Gender < results[j].Gender
				}

				return results[i].Gender > results[j].Gender
			})
		case "height":
			sort.Slice(results, func(i, j int) bool {
				if request.Sort.IsAsc {
					return results[i].Height < results[j].Height
				}

				return results[i].Height > results[j].Height
			})
		}
	}

	return results, totalCount, totalHeight, nil
}
