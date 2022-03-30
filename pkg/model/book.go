package model

	

type Book struct { 
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name"`
	Author string `json:"author"`
	Price float64  `json:"price"`
}
