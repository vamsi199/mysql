package main

import (
	//"fmt"
	"github.com/jinzhu/gorm"
        _ "github.com/go-sql-driver/mysql"
)

func main()  {
	db,err := gorm.Open("mysql", "root:pass@(107.21.2.246:3306)/vamsigo")
	if err!= nil {
		panic(err.Error())
	}
	defer db.Close()
	db.CreateTable(&Owner{})




}
type Owner struct {
	Firstname string
	Lastname string
}



