package config

import "github.com/jinzhu/gorm"

var(
	db *gorm.DB
)

func Connect(){
	conn,err:= gorm.Open("mysql","root:password123@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDB() *gorm.DB{
	return db
}