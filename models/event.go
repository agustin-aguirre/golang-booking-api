package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (event Event) Persist() error {
	query := `
	INSERT INTO Events(name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM Events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
