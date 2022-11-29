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
	count int32
}

type NoteStorageI interface {
	Create(note *Note) (*Note, error)
	// Get(id int64) (*User, error)
	// GetAll(params *GetAllNotesParams) (*GetAllNotesResults, error)
	// Update(notes *Note) error
	// Delete(id int64) error
}