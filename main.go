package main

import (
	"fmt"
	"web/db"
	"web/router"
)

func init(){
	fmt.Println("This is the init method")
	conn := db.DbConnection()
	db.CreatTable(conn)
}

func main() {
	router.Router()
}


