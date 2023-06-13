package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 打开数据库
	db, err := sql.Open("mysql", "root:Zhang890@tcp(127.0.0.1:3306)/dtstack?charset=utf8")
	if err != nil {
		log.Fatalf("open database error: %v\n", err)
	}

	fmt.Println(db)
}
