package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Note structure
type Note struct {
	Content    string    `json:"content" binding:"required"`
	Created_at time.Time `json:"created_at"`
}

type Msg struct {
	Message string `json:"message"`
}

// DB info
const (
	DBHost    = "db"
	DBUser    = "postrgres-user"
	DBPass    = "passme123"
	DBName    = "notes"
	Migration = `CREATE TABLE IF NOT EXISTS notes (
		id serial PRIMARY KEY,
		content text NOT NULL,
		created_at timestamp with time zone DEFAULT current_timestamp
	)`
)

func getNotes() ([]Note, error) {
	const query = "SELECT Content, created_at FROM notes ORDER BY created_at"
	return nil, nil
}

func postNote(note Note) error {
	return nil
}

func main() {
	//fmt.Println("vim-go")
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		results, err := getNotes()

		if err != nil {
			msg := Msg{
				Message: "Unable to get notes",
			}
			json.NewEncoder(w).Encode(msg)
		}

		json.NewEncoder(w).Encode(results)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

	})
}
