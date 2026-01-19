package repository

import (
	"GolangDBMySQL/entity"
	"context"
	"database/sql"
	"errors"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	scriptSql := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, scriptSql, comment.Email, comment.Comment)	
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

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	scriptSql := "SELECT id, email, comment FROM comments WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, scriptSql, id)
	comment := entity.Comment{}
	if err != nil {
		return entity.Comment{}, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("comment not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	scriptSql := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, scriptSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}