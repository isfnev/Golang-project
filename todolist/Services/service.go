package service

import (
	model "Todo/Models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func CreateNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var note model.NoteForDb
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	sqlStatement := `insert into todo (name) values ($1) returning id`
	var insertId int

	db := createConnection()
	defer db.Close()
	err = db.QueryRow(sqlStatement, note.Data).Scan(&insertId)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(insertId)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var note model.NoteInDb
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	sqlStatement := `update todo set name = $1 where id = $2`
	db := createConnection()
	defer db.Close()

	db.Exec(sqlStatement, note.Data, note.TodoId)
	json.NewEncoder(w).Encode("Success")
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sqlStatement := `select * from todo`

	db := createConnection()
	defer db.Close()

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var notes []model.NoteInDb
	var note model.NoteInDb

	for rows.Next() {

		err = rows.Scan(&note.TodoId, &note.Data)

		if err != nil {
			panic(err)
		}
		notes = append(notes, note)

	}
	json.NewEncoder(w).Encode(notes)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var note model.NoteInDb
	json.NewDecoder(r.Body).Decode(&note)

	sqlStatement := `delete from todo where id = $1`
	db := createConnection()
	defer db.Close()

	_, err := db.Exec(sqlStatement, note.TodoId)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(note.TodoId)
}
