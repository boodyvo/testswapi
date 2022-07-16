//go:build local
// +build local

package redis

import (
	"context"
	"fmt"
	"github.com/boodyvo/testswapi/fetcher"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchMovies(t *testing.T) {
	movies, err := fetchMovies(context.Background())
	require.NoError(t, err)
	require.Len(t, movies, 6)
}

func TestFetchCharacters(t *testing.T) {
	movie, err := fetchMovie(context.Background(), "1")
	require.NoError(t, err)
	fmt.Printf("%+v\n", movie)

	characters, err := fetchCharacter(context.Background(), movie.Characters)
	require.NoError(t, err)
	fmt.Println(characters)
}

func TestRedis(t *testing.T) {
	red, err := New("localhost:6379")
	require.NoError(t, err)

	characters, _, _, err := red.ListCharacters(context.Background(), fetcher.CharacterRequest{MovieID: "1"})
	require.NoError(t, err)
	fmt.Println(characters)

	characters, _, _, err = red.ListCharacters(context.Background(), fetcher.CharacterRequest{MovieID: "1"})
	require.NoError(t, err)
	fmt.Println(characters)
}
