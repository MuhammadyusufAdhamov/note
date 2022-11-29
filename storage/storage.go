package storage

import (
	"github.com/MuhammadyusufAdhamov/note/storage/postgres"
	"github.com/MuhammadyusufAdhamov/note/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface{
	User() repo.UserStorageI
	Note() repo.NoteStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
	noteRepo repo.NoteStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUser(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Note() repo.NoteStorageI {
	return s.noteRepo
}