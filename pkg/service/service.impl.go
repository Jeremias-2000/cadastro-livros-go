package service

import (
	"context"

	"github.com/book-management/pkg/model"
	"github.com/jinzhu/gorm"
)

type BookServiceImpl struct {
	db *gorm.DB
	ctx        context.Context
}



func 	NewBookServiceImpl	(db *gorm.DB,ctx *context.Context) BookService {
	return &BookServiceImpl{	
		db: db,
		ctx: *ctx,
	}
}

// GetBooks implements BookService
func (b *BookServiceImpl) GetBooks() []*model.Book {
	var books []*model.Book
	b.db.Find(&books)
	return books
}

// GetProductById implements BookService
func (b *BookServiceImpl) GetBookById(ID *int64) *model.Book {
	var book *model.Book
	b.db.Where("ID=?",ID).Find(&book)
	return book

}


// CreateBook implements BookService
func (b *BookServiceImpl) CreateBook(book *model.Book) (*model.Book,error) {
	b.db.NewRecord(book)
	error := b.db.Create(&book).Error
	if error != nil {
		return nil,error
	}
	return book,nil	
}


// DeleteBook implements BookService
func (b *BookServiceImpl) DeleteBook(ID *int64) *model.Book {
	var book *model.Book
	b.db.Where("ID=?",ID).Delete(&book)
	return book
}





// UpdateBook implements BookService
func (*BookServiceImpl) UpdateBook(*model.Book) error {
	panic("unimplemented")
}

