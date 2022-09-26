package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Note structure
type Note struct {
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// DB info
const (
	DBHost    = "127.0.0.1"
	DBUser    = "gouser"
	DBPass    = "gomysql"
	DBName    = "go_notes"
	Migration = `CREATE TABLE IF NOT EXISTS notes (
		id serial PRIMARY KEY,
		content VARCHAR(150) NOT NULL,
		created_at TIMESTAMP
	)`
)

var db *sql.DB

func getNotes() ([]Note, error) {
	const query = "SELECT Content, created_at FROM notes ORDER BY created_at;"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make([]Note, 0)
	log.Println(results)

	for rows.Next() {
		var content string
		var createdAt time.Time

		err = rows.Scan(&content, &createdAt)

		if err != nil {
			return nil, err
		}

		results = append(results, Note{content, createdAt})
	}

	log.Println(results)
	return results, nil
}

func postNote(note Note) error {
	const query = "INSERT INTO notes(content, created_at) VALUES(?, ?);"
	inserQ, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = inserQ.Exec(note.Content, note.CreatedAt)
	return err
}

func main() {
	var err error
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", DBUser, DBPass, DBHost, DBName)
	db, err = sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// exec migration query
	_, err = db.Query(Migration)
	if err != nil {
		log.Fatal("Error during migration, err = " + err.Error())
		return
	}

	r := gin.Default()

	// fetch all notes
	r.GET("/all", func(ctx *gin.Context) {
		results, err := getNotes()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured: " + err.Error()})
		}
		ctx.JSON(http.StatusOK, results)
	})

	// add/publish note
	r.POST("/add", func(ctx *gin.Context) {
		var note Note
		if ctx.Bind(&note) == nil {
			note.CreatedAt = time.Now().UTC()

			if err := postNote(note); err != nil {

				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured: " + err.Error()})
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "Note successfully added"})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "POST body improper"})
		}
	})

	// Start server
	log.Println("Server running...")
	r.Run(":8000")
}
