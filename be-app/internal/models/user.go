package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Points    int    `json:"points" db:"points"`
}

type UserModel interface {
	CreateUser(user User) error
	GetuserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type UserModelImpl struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModelImpl {
	return &UserModelImpl{DB: db}
}

func (userModel *UserModelImpl) CreateUser(user User) error {
	query := `INSERT INTO Users (email, password, first_name, last_name, points) VALUES (?, ?, ?, ?, ?)`

	exists, err := userModel.userExists(user.Email)

	if err != nil {
		fmt.Println("Error checking if user exists", err)
		return err
	}

	if exists {
		fmt.Println("User already exists")
		return errors.New("User already exists")
	}

	_, err = userModel.DB.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, user.Points)

	if err != nil {
		fmt.Println("Error inserting into users table", err)
		return err
	}

	return nil
}

func (UserModel *UserModelImpl) GetAllUsers() ([]User, error) {
	query := `SELECT * FROM Users`

	rows, err := UserModel.DB.Query(query)

	if err != nil {
		fmt.Println("Error querying users table", err)
		return nil, err
	}

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Points)

		if err != nil {
			fmt.Println("Error scanning user into struct", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userModel *UserModelImpl) GetUserByID(id int) (*User, error) {
	query := `SELECT * FROM Users WHERE id = ? `

	var user User

	row := userModel.DB.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Points)

	if err != nil {
		fmt.Println("Error scanning user into struct", err)
		return nil, err
	}

	return &user, nil
}

func (userModel *UserModelImpl) GetUserByEmail(email string) (*User, error) {
	query := `SELECT * FROM Users WHERE email = ? `

	var user User

	row := userModel.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Points)

	if err != nil {
		fmt.Println("Error scanning user into struct", err)
		return nil, err
	}

	return &user, nil
}

func (userModel *UserModelImpl) UpdateUser(user User) error {
	query := `UPDATE Users SET email = ?, password = ?, first_name = ?, last_name = ?, points = ? WHERE id = ?`

	_, err := userModel.DB.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, user.Points, user.ID)

	if err != nil {
		fmt.Println("Error updating user", err)
		return err
	}

	return nil
}

func (uModel *UserModelImpl) userExists(email string) (bool, error) {
	query := `SELECT COUNT(*) FROM Users WHERE email = ?`

	var count int
	err := uModel.DB.QueryRow(query, email).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
