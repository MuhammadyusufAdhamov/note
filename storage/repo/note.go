package repo

import "time"

type Note struct {
	ID int64
	UserId int
	Title string
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type GetAllNotesParams struct {
	Limit int32
	Page int32
	Search string
}

type GetAllNotesResults struct {
	Notes []*Note
	Count int32
}

type NoteStorageI interface {
	CreateNote(note *Note) (*Note, error)
	GetNote(id int64) (*Note, error)
	GetAllNotes(params *GetAllNotesParams) (*GetAllNotesResults, error)
	UpdateNote(note *Note) (*Note, error)
	DeleteNote(note *Note) (*Note, error)
}