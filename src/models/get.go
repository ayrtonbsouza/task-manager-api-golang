package models

import (
	"task-manager-api/src/db"

	"github.com/google/uuid"
)

func Get(id uuid.UUID) (task Task, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM Task WHERE id = $1", id)

	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Deadline)

	return
}
