package models

import "time"

type Note struct {
	ID int64				`json:"id"`
	UserId int				`json:"user_id"`
	Title string			`json:"title"`
	Description string		`json:"description"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
	DeletedAt time.Time		`json:"deleted_at"`
}

type CreateNoteRequest struct {
	UserId int				`json:"user_id"`
	Title string			`json:"title"`
	Description string		`json:"description"`
}

type GetAllNotesResponse struct {
	Notes []*Note  `json:"note"`
	Count int32    `json:"count"`
}