package model

type NoteInDb struct {
	TodoId int `json:"todoid"`
	Data string `json:"note"`
}

type NoteForDb struct{
	Data string `json:"note"`
}