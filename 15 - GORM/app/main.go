package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "fastbyte:mysqlpass@tcp(localhost:3306)/gorm_test?charset=utf8mb4"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type GormTestModel struct {
	Username string
	Password string
}

type ErrorMsg struct {
	Msg string
}

func main() {
	if err != nil {
		log.Fatal(err)
	}
	cleanup()
	// defer db.Close() // Errors out
	requests()
}

func cleanup() {
	// Remove table if exists
	db.Migrator().DropTable(&GormTestModel{})
	db.Migrator().CreateTable(&GormTestModel{})
	if db.Migrator().HasTable(&GormTestModel{}) {
		fmt.Println("Table created!")
	}
}

func requests() {
	http.HandleFunc("/insert", InsertData)

	fmt.Println("Starting a server: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func InsertData(response http.ResponseWriter, r *http.Request) {
	uname := r.URL.Query().Get("uname")
	pass := r.URL.Query().Get("pass")

	if len(uname) == 0 || len(pass) == 0 {
		ErrorMsg := ErrorMsg{Msg: "Failed to specify uname and passwd in URL."}
		json.NewEncoder(response).Encode(ErrorMsg)
		return
	}

	gormTestModel := GormTestModel{
		Username: uname,
		Password: pass,
	}

	fmt.Println("Inserting data into Table")

	if err = db.Create(&gormTestModel).Error; err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(response).Encode(gormTestModel)
	fmt.Println("Added: ", gormTestModel)
}
