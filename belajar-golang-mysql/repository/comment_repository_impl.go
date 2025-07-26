package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CommentRepositoryImpl struct{
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert (ctx context.Context,comment entity.Comments)(entity.Comments,error){
	script := "INSERT INTO comments(email,comment) VALUES (?,?)"
	result,err := repository.DB.ExecContext(ctx,script,comment.Email,comment.Comment)

	if err != nil {
		return comment,err
	}

	id,err := result.LastInsertId()
	if err != nil{
		return comment,err
	}

	comment.Id = int32(id)
	return comment,err
	
}

func (repository *CommentRepositoryImpl)FindById(ctx context.Context,id int32)(entity.Comments,error){
	script := "SELECT id,email,comment FROM comments where id = ? LIMIT 1"
	rows,err := repository.DB.QueryContext(ctx,script,id)
	defer rows.Close()
	comment := entity.Comments{}
	if err != nil{
		return comment,err
	}
	if rows.Next(){
		rows.Scan(&comment.Id,&comment.Email,&comment.Comment)
		return comment,nil

	}else{
		return comment,errors.New("comment tidak ditemukan")
	}
}

func (repository *CommentRepositoryImpl)FindAll(ctx context.Context)([]entity.Comments,error){
	script := "SELECT id,email,comment FROM comments"
	rows,err := repository.DB.QueryContext(ctx,script)
	defer rows.Close()
	comments := []entity.Comments{}
	if err != nil{
		return nil,err
	}
	for rows.Next(){
		comment := entity.Comments{}
		rows.Scan(&comment.Id,&comment.Email,&comment.Comment)
		comments = append(comments,comment)
	}
	return comments, nil
}

func (repository *CommentRepositoryImpl) DeleteById (ctx context.Context,id int32) (string,error){
	script := "DELETE FROM comments where id = ?"
	result,_ := repository.DB.ExecContext(ctx,script,id)
	rows,err := result.RowsAffected()
	if err != nil{
		return "",err
	}
	if rows == 0{
		return "komentar tidak ditemukan",err
	}
	return "berhasil hapus komentar",nil
}

func (repository *CommentRepositoryImpl)UpdateById(ctx context.Context,id int32,comment entity.Comments)(entity.Comments,error){
	script := "UPDATE comments SET email = ?, comment = ? WHERE id = ?"
	result,err:= repository.DB.ExecContext(ctx,script,comment.Email,comment.Comment,id)
	if err != nil{
		return comment,err
	}
	rows,err := result.RowsAffected()
	if err != nil {
		return comment, err
	}
	if rows == 0 {
	return comment, fmt.Errorf("tidak ditemukan comment dengan id = %d", id)
}
	comment.Id = id
	return comment, nil
}