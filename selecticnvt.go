// test project main.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test:test@tcp(192.168.0.5:3306)/world?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select icntv_id,software_code from epg_device where mac='1819d1068893'")
	CheckErr(err)
	for rows.Next() {
		var id string
		var nick string
		err = rows.Scan(&id, &nick)
		CheckErr(err)
		fmt.Println("id:", id)
		fmt.Println("software_code:", nick)
	}
	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
