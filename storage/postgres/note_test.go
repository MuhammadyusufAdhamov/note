package postgres_test

import (
	"testing"
	"time"

	"github.com/MuhammadyusufAdhamov/note/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createNote(t *testing.T) *repo.Note {
	note, err := strg.Note().CreateNote(&repo.Note{
		UserId: 1,
		Title: faker.Sentence(),
		Description: faker.Sentence(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, note)

	return note
}

func TestGetNote(t *testing.T) {
	c := createNote(t)

	note, err := strg.Note().GetNote(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, note)
}

func TestCreateNote(t *testing.T) {
	createNote(t)
}

func TestUpdateNote(t *testing.T) {
	note, err := strg.Note().UpdateNote(&repo.Note{
		UserId: 1,
		Title: faker.Sentence(),
		Description: faker.Sentence(),
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, note)

}

func TestDeleteNote(t *testing.T) {
	user, err := strg.Note().DeleteNote(&repo.Note{
		DeletedAt: time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestGetAllNotes(t *testing.T) {
	user, err := strg.Note().GetAllNotes(&repo.GetAllNotesParams{
		Limit: 5,
		Page: 1,
		Search: "ab",
	})

	require.NoError(t, err)
	require.NotEmpty(t, user)
}
