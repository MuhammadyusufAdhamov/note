package postgres_test

import (
	"testing"
	"time"

	"github.com/MuhammadyusufAdhamov/note/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	user, err := strg.User().CreateUser(&repo.User{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email: faker.Email(),
		ImageUrl: faker.URL(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestGetUser(t *testing.T) {
	c := createUser(t)

	user, err := strg.User().GetUser(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestUpdateUser(t *testing.T) {

	user, err := strg.User().UpdateUser(&repo.User{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email: faker.Email(),
		ImageUrl: faker.URL(),
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestDeleteUser(t *testing.T) {
	user, err := strg.User().DeleteUser(&repo.User{
		DeletedAt: time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestGetAllUsers(t *testing.T) {
	user, err := strg.User().GetAllUsers(&repo.GetAllUsersParams{
		Limit: 5,
		Page: 1,
		Search: "ab",
	})

	require.NoError(t, err)
	require.NotEmpty(t, user)
}

