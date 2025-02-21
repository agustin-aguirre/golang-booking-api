package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Persist() error {
	query := `
	INSERT INTO Users(email, password)
	VALUES (?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPsw, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPsw)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM Users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var fetchedPassword string
	err := row.Scan(&user.ID, &fetchedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, fetchedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM Users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
