package service

import (
	"github.com/book-management/pkg/model"
	
)

type BookService interface{
	GetBooks()[]*model.Book
	GetBookById(id *int64) *model.Book
	CreateBook(*model.Book) (*model.Book,error)
	UpdateBook(*model.Book) error
	DeleteBook(*int64) *model.Book 
	
}