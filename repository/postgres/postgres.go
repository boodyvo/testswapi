package postgres

import (
	"context"
	"database/sql"
	"github.com/boodyvo/testswapi/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	client *bun.DB
}

func New(dsn string) (*DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	// reset each schema for simplicity of example and testing
	if err := db.ResetModel(context.Background(), (*models.Comment)(nil)); err != nil {
		return nil, err
	}

	return &DB{client: db}, nil
}

func (db *DB) ListCommentsCount(ctx context.Context, movieIDs []string) ([]int, error) {
	var comments []*models.Comment
	// simplification instead of
	// select count(*) c, movie_id c from comments where movie_id in (?);
	results := make([]int, 0, len(movieIDs))
	for _, id := range movieIDs {
		count, err := db.client.NewSelect().Model(&comments).Where("movie_id = ?", id).Count(ctx)
		if err != nil {
			return nil, err
		}

		results = append(results, count)
	}

	return results, nil
}

func (db *DB) ListComments(ctx context.Context, id string) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	if err := db.client.NewSelect().
		Model(&comments).Where("movie_id = ?", id).
		OrderExpr("created_at DESC").
		Scan(ctx); err != nil {
		return nil, err
	}

	return comments, nil
}

func (db *DB) CreateComment(ctx context.Context, comment *models.Comment) error {
	if _, err := db.client.NewInsert().Model(comment).Exec(ctx); err != nil {
		return err
	}

	return nil
}
