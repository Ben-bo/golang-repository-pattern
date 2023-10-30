package repository

import (
	"context"

	"github.com/Ben-bo/golang-repository-pattern/entity"
)

type CommentsRepositoryContract interface {
	insert(ctx context.Context, comment entity.Comments) (entity.Comments, error)
	findAll(ctx context.Context) ([]entity.Comments, error)
	findById(ctx context.Context, id int) (entity.Comments, error)
}
