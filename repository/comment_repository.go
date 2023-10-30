package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ben-bo/golang-repository-pattern/entity"
)

type commentRepository struct {
	DB *sql.DB
}

func NewCommentContract(sql *sql.DB) CommentsRepositoryContract {
	return &commentRepository{DB: sql}
}

func (repo *commentRepository) insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {

	queryDB := "INSERT INTO comments(email,comment) values(?,?)"
	result, err := repo.DB.ExecContext(ctx, queryDB, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	return comment, nil

}

func (repo *commentRepository) findAll(ctx context.Context) ([]entity.Comments, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *commentRepository) findById(ctx context.Context, id int) (entity.Comments, error) {
	queryDB := "SELECT id, email, comment FROM comments WHERE id= ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, queryDB, id)
	comment := entity.Comments{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("comment tidak ditemukan")
	}

}
