package models

import "example.com/rest-api/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Persist() error {
	query := `
	INSERT INTO Users(email, password)
	VALUES (?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(query, user.Email, user.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}
