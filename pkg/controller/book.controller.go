package controller

import (
	
	"fmt"
	"net/http"
	"strconv"

	"github.com/book-management/pkg/model"
	"github.com/book-management/pkg/service"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService service.BookService
}


func New (bookService service.BookService) BookController{
	return BookController{
		bookService: bookService,
	}
}


func (controller *BookController) Routes(r *gin.RouterGroup){
	routes := r.Group("/api")
	routes.GET("/all",controller.GetBooks)
	routes.GET("/find/:id",controller.GetBookById)
	routes.POST("/save",controller.CreateBook)
	routes.PUT("/update",controller.UpdateBook)
	routes.DELETE("/delete/:id",controller	.DeleteBook)		
}

func (b *BookController) GetBooks(ctx *gin.Context){
	
	books := b.bookService.GetBooks()
	ctx.JSON(http.StatusOK,books)
}

func (b *BookController) 	GetBookById(ctx *gin.Context){
	id:= ctx.Param("id")
	bookId,err:= strconv.ParseInt(id,0,0)
	
	if err != nil {
		fmt.Println("error while parsing")
	}
	book:= b.bookService.GetBookById(&bookId)
	
	ctx.JSON(http.StatusOK,book)
}
func (b *BookController) CreateBook(ctx *gin.Context){
	var book model.Book
	if err:= ctx.ShouldBindJSON(&book) ;err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	b.bookService.CreateBook(&book)
	
	ctx.JSON(http.StatusCreated,gin.H{"message":"product registered"})
}
func (b *BookController) UpdateBook(ctx *gin.Context){
	
}
func (b *BookController) DeleteBook(ctx *gin.Context){
	
}
