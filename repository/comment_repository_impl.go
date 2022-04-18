package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entitiy"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entitiy.Comments) (entitiy.Comments, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
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

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entitiy.Comments, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entitiy.Comments{}

	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entitiy.Comments, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entitiy.Comments
	for rows.Next() {
		comment := entitiy.Comments{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
