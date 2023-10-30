package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/Ben-bo/golang-repository-pattern/entity"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentContract(getConnectionDB)

	ctx := context.Background()
	comment := entity.Comments{
		Email:   "test@email.com",
		Comment: "test comment",
	}

	result, err := commentRepository.insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
