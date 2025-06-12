package service

import (
	model "Todo/Models"
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"slices"
)

var notes []model.Note

func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note model.Note
	json.NewDecoder(r.Body).Decode(&note)
	note.ID = rand.IntN(1000000)
	notes = append(notes, note)
	json.NewEncoder(w).Encode(notes)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var note model.Note
	json.NewDecoder(r.Body).Decode(&note)

	for i := range notes {
		if notes[i].ID == note.ID {
			notes[i].Data = note.Data
			break
		}
	}
	json.NewEncoder(w).Encode(notes)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note model.Note
	json.NewDecoder(r.Body).Decode(&note)

	for i := range notes {
		if notes[i].ID == note.ID {
			notes = slices.Delete(notes, i, i+1)
		}
	}
	json.NewEncoder(w).Encode(notes)
}
