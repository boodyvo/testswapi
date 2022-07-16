//go:build local
// +build local

package postgres

import (
	"context"
	"github.com/boodyvo/testswapi/models"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	dsn = "postgres://postgres:postgres@localhost:5656/postgres?sslmode=disable"
)

func TestPostgres(t *testing.T) {
	db, err := New(dsn)
	require.NoError(t, err)
	err = db.CreateComment(context.Background(), &models.Comment{
		MovieID:   "1",
		Text:      "some test comment",
		CreatorIP: "10.0.0.56",
	})
	require.NoError(t, err)

	comments, err := db.ListComments(context.Background(), "2")
	require.NoError(t, err)
	require.Len(t, comments, 0)

	comments, err = db.ListComments(context.Background(), "1")
	require.NoError(t, err)
	require.Len(t, comments, 1)

	commentCounts, err := db.ListCommentsCount(context.Background(), []string{"1", "2"})
	require.NoError(t, err)
	require.Equal(t, []int{1, 0}, commentCounts)
}
