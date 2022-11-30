package postgres

import (
	"fmt"
	"time"

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

func (ur *noteRepo) CreateNote(note *repo.Note) (*repo.Note, error) {
	query := `
		insert into notes(
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

func (ur *noteRepo) GetNote(id int64) (*repo.Note, error) {
	var result repo.Note
	
	query := `
		select
			id,
			user_id,
			title,
			description,
			created_at
		from notes where id=$1
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

func (ur *noteRepo) GetAllNotes(params *repo.GetAllNotesParams) (*repo.GetAllNotesResults, error) {
	result := repo.GetAllNotesResults{
		Notes: make([]*repo.Note, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" limit %d offset %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(
			` where title ilike '%s' or description '%s'`,
			str, str,
		)
	}

	query := `
		select 
			id,
			user_id,
			title,
			description,
			created_at
		from notes
	` + filter + `
	order by created_at desc
	` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var n repo.Note

		err := rows.Scan(
			&n.ID,
			&n.UserId,
			&n.Title,
			&n.Description,
			&n.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Notes = append(result.Notes, &n)
	}

	queryCount := `select count(1) from notes ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *noteRepo) UpdateNote(note *repo.Note) (*repo.Note, error) {
	var result repo.Note
	UpdatedAt := time.Now()

	query := `
		update notes set
			user_id=$1,
			title=$2,
			description=$3
			updated_at=$4
		where id=$5
		returning id,user_id,title,description,created_at,updated_at
	`

	row := ur.db.QueryRow(
		query,
		note.UserId,
		note.Title,
		note.Description,
		UpdatedAt,
		note.ID,
	)

	err := row.Scan(
		&result.ID,
		&result.UserId,
		&result.Title,
		&result.Description,
		&result.CreatedAt,
		UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *noteRepo) DeleteNote(note *repo.Note) (*repo.Note, error) {
	var result repo.Note
	DeletedAt := time.Now()

	query := `update notes set
				deleted_at=$1
			where id=$2
			returning id, deleted_at`
	
	row := ur.db.QueryRow(
		query,
		DeletedAt,
		note.ID,
	)

	err := row.Scan(
		&result.ID,
		&DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}