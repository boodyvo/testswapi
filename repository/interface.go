package repository

import (
	"context"
	"github.com/boodyvo/testswapi/models"
)

type CommentsRepository interface {
	ListCommentsCount(ctx context.Context, movieIDs []string) ([]int, error)
	ListComments(ctx context.Context, id string) ([]*models.Comment, error)
	CreateComment(ctx context.Context, comment *models.Comment) error
}
