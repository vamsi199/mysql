package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func main() {
	db, err := sql.Open("mysql", "root:pass@(107.21.2.246:3306)/vamsigo")
	if err != nil {
		log.Println(err)
	}
	fmt.Print(&db)
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	row, err := db.Query(" update vamsigo.vam set hello ='updated from go' where id=121;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	res2B, _ := json.Marshal(row)
	fmt.Println(string(res2B))
	fmt.Println(row)

}
