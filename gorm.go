package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"fmt"
	"time"
)

func main() {
	own := Owner{
		Firstname:"vamsi",
		Lastname:"gorapalli",
	}
	bo := Book{
		Name:"hello",
		Publishtime:time.Now(),
		OwnerID:own.ID,
	}
	db, err := gorm.Open("mysql", "root:pass@(52.91.38.6:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.SingularTable(true)
	db.DropTableIfExists(&Owner{},&Book{},&Auther{})
	db.CreateTable(&Owner{},&Book{},&Auther{})
	db.Create(&own)
	db.Create(&bo)
	own1 := Owner{
		Firstname:"vamsi1",
		Lastname:"gorapalli1",
	}
	bo1 := Book{
		Name:"hello1",
		Publishtime:time.Now(),
		OwnerID:own.ID,
	}
	db.Debug().Save(&own1)
	db.Debug().Save(&bo1)

	//db.Create(&own) //create a indexes for the first time
	//own.Firstname = "vamsichanged"
	//db.Debug().Save(&own) //to update indexes
	//db.Delete(&own)

}

type Owner struct {
	gorm.Model
	Firstname string
	Lastname  string
	Books []Book
}
type Book struct {
	gorm.Model
	Name string
	Publishtime time.Time
	OwnerID uint `sql:"index"`
	Authers []Auther `gorm:"many2many:books_authers"`

}
type Auther struct {
	gorm.Model
	Name string

}

/*func (o *Owner) TableName() string { //to change the Table namefrom owner to changed
	return "changed"

}*/
