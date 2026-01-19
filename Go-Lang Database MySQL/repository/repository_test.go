package repository

import (
	"GolangDBMySQL"
	"GolangDBMySQL/entity"
	"context"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T){
	commentRepository := NewCommentRepository(GolangDBMySQL.GetConnection())
	
	ctx:= context.Background()
	comment:= entity.Comment{
		Email: "test@gmail.com",
		Comment: "Hello this is test comment",
	}

	result,err :=commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T){
	commentRepository := NewCommentRepository(GolangDBMySQL.GetConnection())

	ctx:= context.Background()

	result,err :=commentRepository.FindById(ctx, 90)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T){
	commentRepository := NewCommentRepository(GolangDBMySQL.GetConnection())

	ctx:= context.Background()
	results,err :=commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}	
}

