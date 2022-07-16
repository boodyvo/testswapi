package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/boodyvo/testswapi/fetcher"
	"github.com/boodyvo/testswapi/models"
)

func TestProcessCharacter(t *testing.T) {
	type testCase struct {
		name               string
		characters         []*models.Character
		request            fetcher.CharacterRequest
		expectedCharacters []*models.Character
		expectedCount      int
		expectedHeight     int
	}

	testCases := []testCase{
		{
			name: "no sort and filter",
			characters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 2", Gender: "male", Height: 200},
				{Name: "test 3", Gender: "female", Height: 150},
			},
			request: fetcher.CharacterRequest{MovieID: "1"},
			expectedCharacters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 2", Gender: "male", Height: 200},
				{Name: "test 3", Gender: "female", Height: 150},
			},
			expectedHeight: 450,
			expectedCount:  3,
		},
		{
			name: "with asc sort",
			characters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 2", Gender: "male", Height: 200},
				{Name: "test 3", Gender: "female", Height: 150},
			},
			request: fetcher.CharacterRequest{
				MovieID: "1",
				Sort: &fetcher.CharacterSort{
					IsAsc:    true,
					SortName: "height",
				},
			},
			expectedCharacters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 3", Gender: "female", Height: 150},
				{Name: "test 2", Gender: "male", Height: 200},
			},
			expectedHeight: 450,
			expectedCount:  3,
		},
		{
			name: "with filter and sort",
			characters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 2", Gender: "male", Height: 200},
				{Name: "test 3", Gender: "female", Height: 150},
			},
			request: fetcher.CharacterRequest{
				MovieID: "1",
				Sort: &fetcher.CharacterSort{
					IsAsc:    false,
					SortName: "height",
				},
				Filter: &fetcher.CharacterFilter{
					Gender: "male",
				},
			},
			expectedCharacters: []*models.Character{
				{Name: "test 2", Gender: "male", Height: 200},
				{Name: "test 1", Gender: "male", Height: 100},
			},
			expectedHeight: 300,
			expectedCount:  2,
		},
		{
			name: "with filter",
			characters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 2", Gender: "male", Height: 200},
				{Name: "test 3", Gender: "female", Height: 150},
			},
			request: fetcher.CharacterRequest{
				MovieID: "1",
				Filter: &fetcher.CharacterFilter{
					Gender: "male",
				},
			},
			expectedCharacters: []*models.Character{
				{Name: "test 1", Gender: "male", Height: 100},
				{Name: "test 2", Gender: "male", Height: 200},
			},
			expectedHeight: 300,
			expectedCount:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			characters, count, height, err := processCharacters(context.Background(), tc.request, tc.characters)
			require.NoError(t, err)
			require.Equal(t, tc.expectedCharacters, characters)
			require.Equal(t, tc.expectedCount, count)
			require.Equal(t, tc.expectedHeight, height)
		})
	}
}
