package main

import (
	"TortoiseAssignment/router"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "jit:pass@tcp(localhost:3306)/TTDB?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}
func main() {

	initDB()
	defer db.Close()
	router := router.Router(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
