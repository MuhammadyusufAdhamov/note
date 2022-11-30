package postgres

import (
	"fmt"
	"time"

	"github.com/MuhammadyusufAdhamov/note/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) CreateUser(user *repo.User) (*repo.User, error) {
	query := `
		insert into users (
			first_name,
			last_name,
			phone_number,
			email,
			image_url
		) values ($1,$2,$3,$4,$5)
		returning id,created_at
	`

	row := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.ImageUrl,
	)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (ur *userRepo) GetUser(id int64) (*repo.User, error) {
	var result repo.User
	
	query := `
		select
			id,
			first_name,
			last_name,
			phone_number,
			email,
			image_url,
			created_at
		from users where id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.ImageUrl,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) GetAllUsers(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	result := repo.GetAllUsersResult{
		Users: make([]*repo.User, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" limit %d offset %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(
			` where first_name ilike '%s' or last_name ilike '%s' or email ilike '%s' `,
			str, str,str,
		)
	}

	query := `
		select 
			id,
			first_name,
			last_name,
			phone_number,
			email,
			image_url,
			created_at
		from users
	` + filter + ` 	
	order by created_at desc
	` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.User

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.PhoneNumber,
			&u.Email,
			&u.ImageUrl,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Users = append(result.Users, &u)
	}

	queryCount := `select count(1) from users ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


func (ur *userRepo) UpdateUser(user *repo.User) (*repo.User, error) {
	var result repo.User
	UpdatedAt := time.Now()

	query := `
		update users set
			first_name=$1,
			last_name=$2,
			phone_number=$3,
			email=$4,
			image_url=$5,
			updated_at=$6
		where id=$7
		returning id,first_name,last_name,phone_number,email,image_url,created_at,updated_at
	`

	row := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.ImageUrl,
		UpdatedAt,
		user.ID,
	)

	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.ImageUrl,
		&result.CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) DeleteUser(user *repo.User) (*repo.User, error) {
	var result repo.User
	DeletedAt := time.Now()

	query := `update users set
				deleted_at=$1
			where id=$2
			returning id, deleted_at`
	
	row := ur.db.QueryRow(
		query,
		DeletedAt,
		user.ID,
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