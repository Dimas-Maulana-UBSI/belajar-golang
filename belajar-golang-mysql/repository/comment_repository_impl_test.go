package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T){
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comments{
		Email: "coba@gmail.com",
		Comment: "halo satu dua tiga",
	}
	result,err := commentRepository.Insert(ctx,comment)
	if err != nil{
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T){
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	result,err := CommentRepository.FindById(ctx,20)
	if err != nil{
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T){
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	result,err := CommentRepository.FindAll(ctx)
	if err != nil{
		panic(err)
	}
	for _,data := range result {
		fmt.Println(data)
	}
	
}

func TestDeleteById(t *testing.T){
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	result,err := CommentRepository.DeleteById(ctx,22)

	if err != nil{
		panic(err)
	}
	fmt.Println(result)
}
func TestUpdateById(t *testing.T){
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comments{
		Email: "emailbaru@gmail.com",
		Comment: "comment baru dari update tes",
	}
	result,err := CommentRepository.UpdateById(ctx,20,comment)

	if err != nil{
		panic(err)
	}
	fmt.Println(result)
}