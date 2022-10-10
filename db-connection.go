package main
 
import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)
 
var Database *gorm.DB
var urlDSN = "root:@tcp(localhost:3306)/blog_post_golang?parseTime=true"
var err error
 
func DataMigration()  {
   
    Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
    if err != nil {
        fmt.Print(err.Error())
        panic("Connection error")
    }
    Database.AutoMigrate(&Employee{})
 
}
