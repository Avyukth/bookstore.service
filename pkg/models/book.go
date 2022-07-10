package models


import(
  "gorm.io/gorm"
  "github.com/avyukth/bookstore.service/pkg/config"
  )

var db *gorm.DB

type Book struct{
  gorm.model
  Name string `gorm:""json:"name"`
  Author string  `json:"author"`
  Publication string  `json:"publication"`
}

func init(){
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}

fun (b, *Book) CreateBook() *Book{
  db.NewRecord(b)
  db.Create(&b)
  return b
}

fun GetAllBooks() []Book{
  var Book []Book
  db.Find(&Books)
  return Books
}


fun GetBookById(Id int64) (*Book, *gorm.DB){
  var getBook Book
  db:= db.Where("ID=?",Id).Find(&getBook)
  return &getBook , db
}

fun DeleteBookById(Id int64) Book {
  var book Book
  db.Where("ID=?",Id).Delete(book)
  return book 
}


