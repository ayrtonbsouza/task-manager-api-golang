package models

import (
	"task-manager-api/src/db"

	"github.com/google/uuid"
)

func Insert(task Task) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	sql := `
	INSERT INTO Task (title, description, status, created_at, updated_at, deadline)
	VALUES ($1, $2, "TODO", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $4)
	RETURNING id
	`

	err = conn.QueryRow(sql, task.Title, task.Description, task.Status, task.Deadline).Scan(&id)
	defer conn.Close()
	return
}
