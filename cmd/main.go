package main

import (
	"context"
	"log"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/book-management/pkg/config"
	"github.com/book-management/pkg/controller"
	"github.com/book-management/pkg/model"
	"github.com/book-management/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	server         *gin.Engine
	bookservice    service.BookService
	bookcontroller controller.BookController
	ctx            context.Context
	db             *gorm.DB
	book           model.Book
)


func init() {
	ctx = context.TODO()
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&book)		
	bookservice = service.NewBookServiceImpl(db, &ctx)
	bookcontroller = controller.New(bookservice)
	server = gin.Default()
}

func main() {
	basepath := server.Group("/v1")
	bookcontroller.Routes(basepath)
	log.Fatal(server.Run(":8080"))
}
