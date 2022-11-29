package repo

import "time"

type User struct {
	ID int64
	FirstName string
	LastName string
	PhoneNumber string
	Email string
	ImageUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time 
}

type GetAllUsersParams struct {
	Limit int32
	Page int32
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type UserStorageI interface {
	CreateUser(u *User) (*User, error)
	GetUser(id int64) (*User, error)
	GetAllUsers(params *GetAllUsersParams) (*GetAllUsersResult, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(u *User) (*User, error)
}
