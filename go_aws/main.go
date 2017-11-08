package main

import ("database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"io"
	//"net/http"

)

func main () {
	db,err := sql.Open("mysql", "root:asAS1234@tcp(godb.c6wigtzq1tu0.us-east-1.rds.amazonaws.com:3306)/go_db")
	if err != nil {
		fmt.Printf( "cannot connect to the database %s",err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("cannot ping %s",err)
	}


}
