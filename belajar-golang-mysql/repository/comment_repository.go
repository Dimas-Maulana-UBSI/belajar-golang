package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type CommentRepository interface{
	Insert(ctx context.Context, comment entity.Comments)(entity.Comments,error)
	FindById(ctx context.Context, id int32)(entity.Comments,error)
	FindAll(ctx context.Context)([]entity.Comments,error)
	DeleteById(ctx context.Context,id int32)(string,error)
	UpdateById(ctx context.Context,id int32, comment entity.Comments)(entity.Comments,error)
}