package postgres

import (
	"github.com/MuhammadyusufAdhamov/note/storage/repo"
	"github.com/jmoiron/sqlx"
)

type noteRepo struct {
	db *sqlx.DB
}

func NewNote(db *sqlx.DB) repo.NoteStorageI {
	return &noteRepo{
		db: db,
	}
}

func (ur *noteRepo) Create(note *repo.Note) (*repo.Note, error) {
	query := `
		insert into note(
			user_id,
			title,
			description
		) values ($1,$2,$3)
		returning id, created_at
	`

	row := ur.db.QueryRow(
		query,
		note.UserId,
		note.Title,
		note.Description,
	)

	err := row.Scan(
		&note.ID,
		&note.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (ur *noteRepo) Get(id int64) (*repo.Note, error) {
	var result repo.Note
	
	query := `
		select
			id,
			user_id,
			title,
			description,
			created_at
		from note where id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.UserId,
		&result.Title,
		&result.Description,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// func (ur *noteRepo) Update(note *repo.Note) (*repo.Note, error) {

// }